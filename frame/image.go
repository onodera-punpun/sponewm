package frame

import (
	"image"
	"os"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/logger"
)

type Image struct {
	*xgraphics.Image
}

func New(ximg *xgraphics.Image) *Image {
	return &Image{ximg}
}

func NewImageA(X *xgbutil.XUtil, side string) *Image {
	file, err := os.Open("/home/onodera/.config/sponewm/images/active_" + side + ".png")
	if err != nil {
		logger.Warning.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		logger.Warning.Fatalln(err)
	}

	return New(xgraphics.NewConvert(X, img))
}

func NewImageI(X *xgbutil.XUtil, side string) *Image {
	file, err := os.Open("/home/onodera/.config/sponewm/images/inactive_" + side + ".png")
	if err != nil {
		logger.Warning.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		logger.Warning.Fatalln(err)
	}

	return New(xgraphics.NewConvert(X, img))
}

func ImageSize(side string) int {
	file, err := os.Open("/home/onodera/.config/sponewm/images/active_" + side + ".png")
	if err != nil {
		logger.Warning.Fatalln(err)
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		logger.Warning.Fatalln(err)
	}
	var size int
	if side == "top" || side == "bottom" {
		size = img.Height
	} else if side == "left" || side == "right" {
		size = img.Width
	}

	return size
}
