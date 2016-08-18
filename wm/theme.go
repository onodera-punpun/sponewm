package wm

import (
	"image"
	"os"

	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/frame"
	"github.com/onodera-punpun/sponewm/logger"
)

type ThemeConfig struct {
	decorTopA, decorTopI                 *xgraphics.Image
	decorBottomA, decorBottomI           *xgraphics.Image
	decorLeftA, decorLeftI               *xgraphics.Image
	decorRightA, decorRightI             *xgraphics.Image
	decorTopLeftA, decorTopLeftI         *xgraphics.Image
	decorTopRightA, decorTopRightI       *xgraphics.Image
	decorBottomLeftA, decorBottomLeftI   *xgraphics.Image
	decorBottomRightA, decorBottomRightI *xgraphics.Image
	decorSizeTop                         int
	decorSizeBottom                      int
	decorSizeLeft                        int
	decorSizeRight                       int
}

func (td ThemeConfig) FrameTheme() *frame.DecorTheme {
	return &frame.DecorTheme{
		DecorTopA:         td.decorTopA,
		DecorTopI:         td.decorTopI,
		DecorBottomA:      td.decorBottomA,
		DecorBottomI:      td.decorBottomI,
		DecorLeftA:        td.decorLeftA,
		DecorLeftI:        td.decorLeftI,
		DecorRightA:       td.decorRightA,
		DecorRightI:       td.decorRightI,
		DecorTopLeftA:     td.decorTopLeftA,
		DecorTopLeftI:     td.decorTopLeftI,
		DecorTopRightA:    td.decorTopRightA,
		DecorTopRightI:    td.decorTopRightI,
		DecorBottomLeftA:  td.decorBottomLeftA,
		DecorBottomLeftI:  td.decorBottomLeftI,
		DecorBottomRightA: td.decorBottomRightA,
		DecorBottomRightI: td.decorBottomRightI,
		DecorSizeTop:      td.decorSizeTop,
		DecorSizeBottom:   td.decorSizeBottom,
		DecorSizeLeft:     td.decorSizeLeft,
		DecorSizeRight:    td.decorSizeRight,
	}
}

func newTheme() *ThemeConfig {
	return &ThemeConfig{
		decorTopA:         newImage("active_top").Image,
		decorTopI:         newImage("inactive_top").Image,
		decorBottomA:      newImage("active_bottom").Image,
		decorBottomI:      newImage("inactive_bottom").Image,
		decorLeftA:        newImage("active_left").Image,
		decorLeftI:        newImage("inactive_left").Image,
		decorRightA:       newImage("active_right").Image,
		decorRightI:       newImage("inactive_right").Image,
		decorTopLeftA:     newImage("active_topleft").Image,
		decorTopLeftI:     newImage("inactive_topleft").Image,
		decorTopRightA:    newImage("active_topright").Image,
		decorTopRightI:    newImage("inactive_topright").Image,
		decorBottomLeftA:  newImage("active_bottomleft").Image,
		decorBottomLeftI:  newImage("inactive_bottomleft").Image,
		decorBottomRightA: newImage("active_bottomright").Image,
		decorBottomRightI: newImage("inactive_bottomright").Image,
		decorSizeTop:      imageSize("top"),
		decorSizeBottom:   imageSize("bottom"),
		decorSizeLeft:     imageSize("left"),
		decorSizeRight:    imageSize("right"),
	}
}

func loadTheme() (*ThemeConfig) {
	return newTheme()
}

type Image struct {
	*xgraphics.Image
}

func New(ximg *xgraphics.Image) *Image {
	return &Image{ximg}
}

func newImage(side string) *Image {
	file, err := os.Open("/home/onodera/.config/sponewm/images/" + side + ".png")
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
