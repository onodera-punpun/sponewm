package misc

import (
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
