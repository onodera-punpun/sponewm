package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"runtime/pprof"
	"strings"
	"syscall"
	"time"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/mousebind"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/onodera-punpun/sponewm/commands"
	"github.com/onodera-punpun/sponewm/cursors"
	"github.com/onodera-punpun/sponewm/event"
	"github.com/onodera-punpun/sponewm/frame"

	"github.com/onodera-punpun/sponewm/focus"
	"github.com/onodera-punpun/sponewm/logger"
	"github.com/onodera-punpun/sponewm/misc"
	"github.com/onodera-punpun/sponewm/stack"
	"github.com/onodera-punpun/sponewm/wm"
	"github.com/onodera-punpun/sponewm/xclient"
)

var (
	flagGoMaxProcs     = runtime.NumCPU()
	flagLogLevel       = 2
	flagLogColors      = false
	flagReplace        = false
	flagConfigDir      = ""
	flagCpuProfile     = ""
	flagSponeRestarted = false
	flagShowSocket     = false
)

func init() {
	flag.IntVar(&flagGoMaxProcs, "p", flagGoMaxProcs,
		"The maximum number of CPUs that can be executing simultaneously.")
	flag.IntVar(&flagLogLevel, "log-level", flagLogLevel,
		"The logging level of SponeWM. Valid values are 0, 1, 2, 3 or 4.\n"+
			"Higher numbers result in SponeWM producing more output.")
	flag.BoolVar(&flagLogColors, "log-colors", flagLogColors,
		"Whether to output logging data with terminal colors.")
	flag.BoolVar(&flagReplace, "replace", flagReplace,
		"When set, SponeWM will attempt to replace a currently running\n"+
			"window manager. If this is not set, and another window manager\n"+
			"is running, SponeWM will exit.")
	flag.StringVar(&flagConfigDir, "config-dir", flagConfigDir,
		"Override the location of the configuration files. When this\n"+
			"is not set, the following paths (roughly) will be checked\n"+
			"in order: $XDG_CONFIG_DIR/sponewm, /etc/xdg/sponewm,\n"+
			"$GOPATH/src/github.com/onodera-punpun/sponewm/config")
	flag.BoolVar(&flagSponeRestarted, "spone-restarted", flagSponeRestarted,
		"DO NOT USE. INTERNAL SPONE USE ONLY.")

	flag.BoolVar(&flagShowSocket, "show-socket", flagShowSocket,
		"When set, the command will detect if SponeWM is already running,\n"+
			"and if so, outputs the file path to the current socket.")

	flag.StringVar(&flagCpuProfile, "cpuprofile", flagCpuProfile,
		"When set, a CPU profile will be written to the file specified.")

	flag.Usage = usage
	flag.Parse()

	runtime.GOMAXPROCS(flagGoMaxProcs)
	logger.Colors(flagLogColors)
	logger.LevelSet(flagLogLevel)

	// If the log level is 0, don't show XGB log output either.
	if flagLogLevel == 0 || flagShowSocket {
		xgb.Logger = log.New(ioutil.Discard, "", 0)
	}
}

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		logger.Error.Println(err)
		logger.Error.Fatalln("Error connecting to X, quitting...")
	}
	defer X.Conn().Close()

	if flagShowSocket {
		showSocketPath(X)
		return
	}

	if len(flagConfigDir) > 0 {
		misc.ConfigPaths.Override = flagConfigDir
	}

	keybind.Initialize(X)
	mousebind.Initialize(X)
	focus.Initialize(X)
	stack.Initialize(X)
	cursors.Initialize(X)
	wm.Initialize(X, commands.Env, newHacks())

	// Initialize event handlers on the root window.
	rootInit(X)

	// Tell everyone what we support.
	setSupported()

	// Start up the IPC command listener.
	go ipc(X)

	// And start up the IPC event notifier.
	go event.Notifier(X, socketFilePath(X))

	// Just before starting the main event loop, check to see if there are
	// any clients that already exist that we should manage.
	manageExistingClients()

	// Now make sure that clients are in the appropriate visible state.
	for _, wrk := range wm.Heads.Workspaces.Wrks {
		if wrk.IsVisible() {
			wrk.Show()
		} else {
			wrk.Hide()
		}
	}
	wm.Heads.ApplyStruts(wm.Clients)

	wm.FocusFallback()
	wm.Startup = false
	pingBefore, pingAfter, pingQuit := xevent.MainPing(X)

	if len(flagCpuProfile) > 0 {
		f, err := os.Create(flagCpuProfile)
		if err != nil {
			logger.Error.Fatalf("%s\n", err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

EVENTLOOP:
	for {
		select {
		case <-pingBefore:
			// Wait for the event to finish processing.
			<-pingAfter
		case f := <-commands.SafeExec:
			commands.SafeReturn <- f()
		case <-pingQuit:
			break EVENTLOOP
		}
	}
	if wm.Restart {
		event.Notify(event.Restarting{})
		for _, client := range wm.Clients {
			c := client.(*xclient.Client)

			// TODO: Erm, can I remove this?
			if _, ok := c.Frame().(*frame.Decor); ok {
				c.FrameNada()
			}
		}
		time.Sleep(1 * time.Second)

		// We need to tell the next invocation of SponeWM that it is being
		// *restarted*.
		found := false
		for _, arg := range os.Args {
			if strings.ToLower(strings.TrimSpace(arg)) == "--spone-restarted" {
				found = true
			}
		}
		if !found {
			os.Args = append(os.Args, "--spone-restarted")
		}
		logger.Message.Println("The user has told us to restart...\n\n\n")
		if err := syscall.Exec(os.Args[0], os.Args, os.Environ()); err != nil {
			logger.Error.Fatalf("Could not exec '%s': %s",
				strings.Join(os.Args, " "), err)
		}
	}
}

func setSupported() {
	// See COMPLIANCE for what's supported and what isn't (plus rationale).
	// ewmhSupported comes from ewmh_supported.go and is automatically
	// generated from the COMPLIANCE file.

	// Set supported atoms
	ewmh.SupportedSet(wm.X, ewmhSupported)

	// While we're at it, set the supporting wm hint too.
	ewmh.SupportingWmCheckSet(wm.X, wm.X.RootWin(), wm.X.Dummy())
	ewmh.SupportingWmCheckSet(wm.X, wm.X.Dummy(), wm.X.Dummy())
	ewmh.WmNameSet(wm.X, wm.X.Dummy(), "SponeWM")
}

// manageExistingClients traverse the window tree and tries to manage all
// top-level clients. Clients that are not in the Unmapped state will be
// managed.
func manageExistingClients() {
	tree, err := xproto.QueryTree(wm.X.Conn(), wm.Root.Id).Reply()
	if err != nil {
		logger.Warning.Printf("Could not issue QueryTree request: %s", err)
		return
	}
	for _, potential := range tree.Children {
		// Ignore our own dummy window...
		if potential == wm.X.Dummy() {
			continue
		}

		attrs, err := xproto.GetWindowAttributes(wm.X.Conn(), potential).Reply()
		if err != nil {
			continue
		}
		if attrs.MapState == xproto.MapStateUnmapped {
			continue
		}
		logger.Message.Printf("Managing existing client %d", potential)
		xclient.New(potential)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "\nUsage: %s [flags]\n", path.Base(os.Args[0]))
	flag.VisitAll(func(fg *flag.Flag) {
		// Don't let users know about flags they shouldn't use.
		if fg.Name == "spone-restarted" {
			return
		}
		fmt.Printf("--%s=\"%s\"\n\t%s\n", fg.Name, fg.DefValue,
			strings.Replace(fg.Usage, "\n", "\n\t", -1))
	})
	os.Exit(1)
}

func showSocketPath(X *xgbutil.XUtil) {
	currentWM, err := ewmh.GetEwmhWM(X)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if strings.ToLower(currentWM) != "sponewm" {
		fmt.Fprintf(os.Stderr, "Could not detect a SponeWM instance. "+
			"(Found '%s' instead.)\n", currentWM)
		os.Exit(1)
	}
	fmt.Println(socketFilePath(X))
	os.Exit(0)
}
