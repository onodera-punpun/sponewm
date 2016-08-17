package wm

// TODO: This is a fucking mess, read up on structs
// and shit, and clean this up.

import (
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/onodera-punpun/sponewm/frame"
	"github.com/onodera-punpun/sponewm/misc"
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
		DecorTopA:       td.decorTopA,
		DecorTopI:       td.decorTopI,
		DecorBottomA:    td.decorBottomA,
		DecorBottomI:    td.decorBottomI,
		DecorLeftA:      td.decorLeftA,
		DecorLeftI:      td.decorLeftI,
		DecorRightA:     td.decorRightA,
		DecorRightI:     td.decorRightI,
		DecorSizeTop:    td.decorSizeTop,
		DecorSizeBottom: td.decorSizeBottom,
		DecorSizeLeft:   td.decorSizeLeft,
		DecorSizeRight:  td.decorSizeRight,
	}
}

func newTheme() *ThemeDecor {
	return &ThemeDecor{
		decorTopA:       builtInImage(misc.DecorTopAPng),
		decorTopI:       builtInImage(misc.DecorTopIPng),
		decorBottomA:    builtInImage(misc.DecorBottomAPng),
		decorBottomI:    builtInImage(misc.DecorBottomIPng),
		decorLeftA:      builtInImage(misc.DecorLeftAPng),
		decorLeftI:      builtInImage(misc.DecorLeftIPng),
		decorRightA:     builtInImage(misc.DecorRightAPng),
		decorRightI:     builtInImage(misc.DecorRightIPng),
		decorSizeTop:    20,
		decorSizeBottom: 10,
		decorSizeLeft:   10,
		decorSizeRight:  10,
	}
}

func loadTheme() (*ThemeDecor, error) {
	theme := newTheme()

	xgraphics.NewFileName(X, "active_top")
	xgraphics.NewFileName(X, "inactive_top")
	xgraphics.NewFileName(X, "active_bottom")
	xgraphics.NewFileName(X, "inactive_bottom")
	xgraphics.NewFileName(X, "active_left")
	xgraphics.NewFileName(X, "inactive_left")
	xgraphics.NewFileName(X, "active_right")
	xgraphics.NewFileName(X, "inactive_right")

	return theme, nil
}

func builtInImage(builtInData []byte) *xgraphics.Image {
	img, err := xgraphics.NewBytes(X, builtInData)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	return img
}
