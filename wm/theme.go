package wm

import (
	"image"
	"os"

	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/frame"
	"github.com/onodera-punpun/sponewm/logger"
)

type ThemeDecor struct {
	decorTopA, decorTopI       *xgraphics.Image
	decorBottomA, decorBottomI *xgraphics.Image
	decorLeftA, decorLeftI     *xgraphics.Image
	decorRightA, decorRightI   *xgraphics.Image
	decorSizeTop               int
	decorSizeBottom            int
	decorSizeLeft              int
	decorSizeRight             int
}

func (td ThemeDecor) FrameTheme() *frame.DecorTheme {
	return &frame.DecorTheme{
		DecorTopA:       newImageA("top").Image,
		DecorTopI:       newImageI("top").Image,
		DecorBottomA:    newImageA("bottom").Image,
		DecorBottomI:    newImageI("bottom").Image,
		DecorLeftA:      newImageA("left").Image,
		DecorLeftI:      newImageI("left").Image,
		DecorRightA:     newImageA("right").Image,
		DecorRightI:     newImageI("right").Image,
		DecorSizeTop:    imageSize("top"),
		DecorSizeBottom: imageSize("bottom"),
		DecorSizeLeft:   imageSize("left"),
		DecorSizeRight:  imageSize("right"),
	}
}

type Image struct {
	*xgraphics.Image
}

func New(ximg *xgraphics.Image) *Image {
	return &Image{ximg}
}

func newImageA(side string) *Image {
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

func newImageI(side string) *Image {
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

func imageSize(side string) int {
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
