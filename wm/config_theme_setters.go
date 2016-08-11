package wm

import (
	"strings"

	"github.com/onodera-punpun/wingo/logger"
	"github.com/onodera-punpun/wingo/render"
	"github.com/onodera-punpun/wingo/wini"
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

func setFloat(k wini.Key, place *float64) {
	if v, ok := getLastFloat(k); ok {
		*place = float64(v)
	}
}

func getLastFloat(k wini.Key) (float64, bool) {
	vals, err := k.Floats()
	if err != nil {
		logger.Warning.Println(err)
		return 0.0, false
	} else if len(vals) == 0 {
		logger.Warning.Println(k.Err("No values found."))
		return 0.0, false
	}

	return vals[len(vals)-1], true
}

func setColor(k wini.Key, clr *render.Color) {
	// Check to make sure we have a value for this key
	vals := k.Strings()
	if len(vals) == 0 {
		logger.Warning.Println(k.Err("No values found."))
		return
	}

	// Use the last value
	val := vals[len(vals)-1]

	// TODO: Can I simplify this?
	if strings.Index(val, " ") == -1 {
		if start, ok := getLastInt(k); ok {
			clr.ColorSet(start)
		}
		return
	}
}
