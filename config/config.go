package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/onodera-punpun/sponewm/logger"
)

var val map[string]interface{}
var key map[string]interface{}
var SettingsVal map[string]interface{}
var SettingsKey map[string]interface{}
var BindingsVal map[string]interface{}
var BindingsKey map[string]interface{}

func Initialize() {
	configs := []string{"settings", "bindings"}

	for _, basename := range configs {
		defaults := defaultConfig(basename)
		var parsed map[string]interface{}

		filename := ConfigDir() + "/" + basename + ".json"
		if _, e := os.Stat(filename); e == nil {
			input, err := ioutil.ReadFile(filename)
			if err != nil {
				logger.Error.Fatalf("Error reading"+basename+".json: %s", err)
			}

			err = json.Unmarshal(input, &parsed)
			if err != nil {
				logger.Error.Fatalf("Error reading"+basename+".json: %s", err)
			}
		}

		val = make(map[string]interface{})
		key = make(map[string]interface{})
		for k, v := range defaults {
			val[k] = v
			key[k] = k
		}
		for k, v := range parsed {
			val[k] = v
			key[k] = k
		}

		err := writeConfig(filename)
		if err != nil {
			logger.Error.Fatalf("Error reading"+basename+".json: %s", err)
		}

		if basename == "settings" {
			SettingsKey = key
			SettingsVal = val
		} else {
			BindingsKey = key
			BindingsVal = val
		}
	}
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
		logger.Error.Fatalf("Something is screwy. SponeWM could not detect "+
			"valid values in your XDG_CONFIG_HOME ('%s') or HOME ('%s') "+
			"environment variables.", xdgHome, home)
	}

	return configDir
}

func writeConfig(filename string,) error {
	var err error

	if _, e := os.Stat(ConfigDir()); e == nil {
		txt, _ := json.MarshalIndent(val, "", "    ")
		err = ioutil.WriteFile(filename, txt, 0644)
	}

	return err
}

func defaultConfig(basename string) map[string]interface{} {
	if basename == "settings" {
		return map[string]interface{}{
			"defaultlayout":     "Floating",
			"focusfollowsmouse": true,
			"raisefollowsmouse": false,
			"floatpadding":      40,
			"gap":               20,
			"tilepadding":       80,
			"workspaces": []interface{}{
				"www",
				"irc",
				"src",
			},
		}
	} else {
		return map[string]interface{}{
			"root": map[string]interface{}{
				"1":      "Focus \":mouse:\"",
				"4":      "Workspace (GetWorkspacePrev)",
				"5":      "Workspace (GetWorkspaceNext)",
				"Mod4-4": "Workspace (GetWorkspacePrev)",
				"Mod4-5": "Workspace (GetWorkspacePrev)",
			},
			"client": map[string]interface{}{
				"1":      "Focus \":mouse:\"",
				"4":      "Workspace (GetWorkspacePrev)",
				"5":      "Workspace (GetWorkspaceNext)",
				"Mod4-4": "Workspace (GetWorkspacePrev)",
				"Mod4-5": "Workspace (GetWorkspacePrev)",
			},
		}
	}
}
