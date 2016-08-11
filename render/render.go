package render

import (
	"image"
)

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

const (
	BorderTop = 1 << iota
	BorderRight
	BorderBottom
	BorderLeft
)

const (
	DiagTopLeft = iota
	DiagTopRight
	DiagBottomLeft
	DiagBottomRight
)

type Image struct {
	*xgraphics.Image
}

func New(ximg *xgraphics.Image) *Image {
	return &Image{ximg}
}

func NewSolid(X *xgbutil.XUtil, bgColor Color, width, height int) *Image {
	img := New(xgraphics.New(X, image.Rect(0, 0, width, height)))

	r, g, b := bgColor.RGB8()
	img.ForExp(func(x, y int) (uint8, uint8, uint8, uint8) {
		return r, g, b, 0xff
	})
	return img
}

func NewBorder(X *xgbutil.XUtil, borderType int, borderColor,
	bgColor Color, width, height) *Image {

	img := New(xgraphics.New(X, image.Rect(0, 0, width, height)))

	r, g, b := bgColor.RGB8()
	img.ForExp(func(x, y int) (uint8, uint8, uint8, uint8) {
		return r, g, b, 0xff
	})

	img.ThinBorder(borderType, borderColor)
	return img
}

func NewCorner(X *xgbutil.XUtil, borderType int, borderColor,
	bgColor Color, width, height, diagonal int) *Image {

	return NewBorder(X, borderType, borderColor, bgColor,
			width, height, 0, 0)
}

// XXX: Optimize.
func (img *Image) ThinBorder(borderType int, borderColor Color) {
	borderClr := borderColor.ImageColor()
	width, height := img.Bounds().Dx(), img.Bounds().Dy()

	// Now go through and add a "thin border."
	// It's always one pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if borderType&BorderTop > 0 && y == 0 {
				img.Set(x, y, borderClr)
			}
			if borderType&BorderRight > 0 && x == width-1 {
				img.Set(x, y, borderClr)
			}
			if borderType&BorderBottom > 0 && y == height-1 {
				img.Set(x, y, borderClr)
			}
			if borderType&BorderLeft > 0 && x == 0 {
				img.Set(x, y, borderClr)
			}
		}
	}
}
