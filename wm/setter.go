package wm

import (
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/logger"
	"github.com/onodera-punpun/sponewm/wini"
)

func setString(k wini.Key, place *string) {
	if v, ok := getLastString(k); ok {
		*place = v
	}
}

func getLastString(k wini.Key) (string, bool) {
	vals := k.Strings()
	if len(vals) == 0 {
		logger.Warning.Println(k.Err("No values found."))
		return "", false
	}

	return vals[len(vals)-1], true
}

func setBool(k wini.Key, place *bool) {
	if v, ok := getLastBool(k); ok {
		*place = v
	}
}

func getLastBool(k wini.Key) (bool, bool) {
	vals, err := k.Bools()
	if err != nil {
		logger.Warning.Println(err)
		return false, false
	} else if len(vals) == 0 {
		logger.Warning.Println(k.Err("No values found."))
		return false, false
	}

	return vals[len(vals)-1], true
}

func setInt(k wini.Key, place *int) {
	if v, ok := getLastInt(k); ok {
		*place = int(v)
	}
}

func getLastInt(k wini.Key) (int, bool) {
	vals, err := k.Ints()
	if err != nil {
		logger.Warning.Println(err)
		return 0, false
	} else if len(vals) == 0 {
		logger.Warning.Println(k.Err("No values found."))
		return 0, false
	}

	return vals[len(vals)-1], true
}

func setImage(k wini.Key, place **xgraphics.Image) {
	if v, ok := getLastString(k); ok {
		img, err := xgraphics.NewFileName(X, v)
		if err != nil {
			logger.Warning.Printf(
				"Could not load '%s' as a png image because: %v", v, err)
			return
		}
		*place = img
	}
}
