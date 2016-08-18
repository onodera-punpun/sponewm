package misc

import (
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/xdg"

	"github.com/onodera-punpun/sponewm/logger"
)

var ConfigPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "sponewm",
	GoImportPath: "github.com/onodera-punpun/sponewm/config",
}

func ConfigFile(name string) string {
	fpath, err := ConfigPaths.ConfigFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	return fpath
}

func ConfigDir() string {
	var configDir string

	xdgHome := os.Getenv("XDG_CONFIG_HOME")
	home := os.Getenv("HOME")
	if len(xdgHome) > 0 && strings.HasPrefix(xdgHome, "/") {
		configDir = path.Join(xdgHome, "sponewm")
	} else if len(home) > 0 && strings.HasPrefix(home, "/") {
		configDir = path.Join(home, ".config", "sponewm")
	} else {
		logger.Error.Fatalf("Something is screwy. Wingo could not detect "+
			"valid values in your XDG_CONFIG_HOME ('%s') or HOME ('%s') "+
			"environment variables.", xdgHome, home)
	}

	return configDir
}
