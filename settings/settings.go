package settings

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/onodera-punpun/sponewm/logger"
)

var Settings map[string]interface{}

func ConfigDir() string {
	var configDir string

	xdgHome := os.Getenv("XDG_CONFIG_HOME")
	home := os.Getenv("HOME")
	if len(xdgHome) > 0 && strings.HasPrefix(xdgHome, "/") {
		configDir = path.Join(xdgHome, "sponewm")
	} else if len(home) > 0 && strings.HasPrefix(home, "/") {
		configDir = path.Join(home, ".config", "sponewm")
	} else {
		logger.Error.Fatalf("Something is screwy. SponeWM could not detect "+
			"valid values in your XDG_CONFIG_HOME ('%s') or HOME ('%s') "+
			"environment variables.", xdgHome, home)
	}

	return configDir
}

func Initialize() {
	defaults := defaultSettings()
	var parsed map[string]interface{}

	filename := ConfigDir() + "/settings.json"
	if _, e := os.Stat(filename); e == nil {
		input, err := ioutil.ReadFile(filename)
		if err != nil {
			logger.Error.Fatalf("Error reading settings.json: %s", err)
		}

		err = json.Unmarshal(input, &parsed)
		if err != nil {
			logger.Error.Fatalf("Error reading settings.json: %s", err)
		}
	}

	Settings = make(map[string]interface{})
	for k, v := range defaults {
		Settings[k] = v
	}
	for k, v := range parsed {
		Settings[k] = v
	}

	err := writeSettings(filename)
	if err != nil {
		logger.Error.Fatalf("Error writing settings.json: %s", err)
	}
}

func writeSettings(filename string) error {
	var err error

	if _, e := os.Stat(ConfigDir()); e == nil {
		txt, _ := json.MarshalIndent(Settings, "", "    ")
		err = ioutil.WriteFile(filename, txt, 0644)
	}

	return err
}

func defaultSettings() map[string]interface{} {
	return map[string]interface{}{
		"defaultlayout":     "Floating",
		"focusfollowsmouse": true,
		"raisefollowsmouse": false,
		"floatpadding":      40,
		"gap":               20,
		"tilepadding":       80,
		"shell":             "bash",
		"workspaces":        []string{"www", "irc", "src"},
	}
}
