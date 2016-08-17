package misc

import (
	"io/ioutil"

	"github.com/BurntSushi/xdg"

	"github.com/onodera-punpun/sponewm/logger"
)

var ConfigPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "sponewm",
	GoImportPath: "github.com/onodera-punpun/sponewm/config",
}

var DataPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "sponewm",
	GoImportPath: "github.com/onodera-punpun/sponewm/data",
}

func ConfigFile(name string) string {
	fpath, err := ConfigPaths.ConfigFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	return fpath
}

func DataFile(name string) []byte {
	fpath, err := DataPaths.DataFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		logger.Error.Fatalf("Could not read %s: %s", fpath, err)
	}
	return bs
}
