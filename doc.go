package main

/*
SponeWM is an X window manager written in pure Go that supports floating and
tiling window placement policies. It is mostly EWMH and ICCCM compliant. Its
"unique" features are per-monitor workspaces and support for both floating and
automatic placement policies (where neither is an after thought).

There is more documentation/guides/compliance in the project directory.

Usage:
	spone [flags]

The flags are:
	--replace
		When set, SponeWM will attempt to replace the currently running window
		manager. This does NOT change your default window manager or edit any
		files. The worst thing that can happen is X crashing.
	--config-dir directory
		When set, SponeWM will always try to read configuration files in the
		directory specified first. (SponeWM will otherwise default to
		$HOME/.config/sponewm or /etc/xdg/sponewm.)
	-p num-cpus
		The maximum number of CPUs that can be executing simultaneously.
		By default, this is set to the number of CPUs detected by the Go
		runtime. Anecdotally, SponeWM feels snappier in this case. When debugging
		however, this should be set to '1' in order to see stack traces in
		their entirety if SponeWM crashes.
	--log-level level
		The logging level of SponeWM. Valid values are 0, 1, 2, 3 or 4. Higher
		numbers result in more logging. When running normally, this should be
		set to 2, which includes errors and warnings. When debugging, this
		should be set to 3, which includes messages that usually are emitted
		in certain state transitions.
	--log-colors
		When set, the log output will highlight errors and warning differently
		from other text using terminal escape sequences.
	--cpuprofile prof-file
		When set, a CPU profile will be written to prof-file when SponeWM exits.

*/
