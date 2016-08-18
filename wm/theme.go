package wm

import (
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/frame"
	"github.com/onodera-punpun/sponewm/logger"
	"github.com/onodera-punpun/sponewm/misc"
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
		decorTopA:         newImage("active_top"),
		decorTopI:         newImage("inactive_top"),
		decorBottomA:      newImage("active_bottom"),
		decorBottomI:      newImage("inactive_bottom"),
		decorLeftA:        newImage("active_left"),
		decorLeftI:        newImage("inactive_left"),
		decorRightA:       newImage("active_right"),
		decorRightI:       newImage("inactive_right"),
		decorTopLeftA:     newImage("active_topleft"),
		decorTopLeftI:     newImage("inactive_topleft"),
		decorTopRightA:    newImage("active_topright"),
		decorTopRightI:    newImage("inactive_topright"),
		decorBottomLeftA:  newImage("active_bottomleft"),
		decorBottomLeftI:  newImage("inactive_bottomleft"),
		decorBottomRightA: newImage("active_bottomright"),
		decorBottomRightI: newImage("inactive_bottomright"),
		decorSizeTop:      newImage("active_top").Bounds().Dy(),
		decorSizeBottom:   newImage("active_bottom").Bounds().Dx(),
		decorSizeLeft:     newImage("active_left").Bounds().Dy(),
		decorSizeRight:    newImage("active_right").Bounds().Dy(),
	}
}

func loadTheme() *ThemeConfig {
	return newTheme()
}

type Image struct {
	*xgraphics.Image
}

func New(pix *xgraphics.Image) *Image {
	return &Image{pix}
}

func newImage(side string) *xgraphics.Image {
	pix, err := xgraphics.NewFileName(X, misc.ConfigDir()+"/images/"+side+".png")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	return pix
}
