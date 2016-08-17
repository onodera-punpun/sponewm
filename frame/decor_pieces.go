package frame

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil/xwindow"

	"github.com/onodera-punpun/sponewm/cursors"
	"github.com/onodera-punpun/sponewm/logger"
)

func (f *Decor) newPieceWindow(ident string,
	cursor xproto.Cursor) *xwindow.Window {

	win, err := xwindow.Generate(f.X)
	if err != nil {
		logger.Error.Printf("Could not create a frame window for client "+
			"with id '%d' because: %s", f.client.Id(), err)
		logger.Error.Fatalf("In a state where no new windows can be created. " +
			"Unfortunately, we must exit.")
	}

	err = win.CreateChecked(f.parent.Id, 0, 0, 1, 1,
		xproto.CwBackPixmap|xproto.CwEventMask|xproto.CwCursor,
		xproto.BackPixmapParentRelative,
		xproto.EventMaskButtonPress|xproto.EventMaskButtonRelease|
			xproto.EventMaskButtonMotion|xproto.EventMaskPointerMotion,
		uint32(cursor))
	if err != nil {
		logger.Warning.Println(err)
	}

	f.client.FramePieceMouseSetup("decor_"+ident, win.Id)

	return win
}

func (f *Decor) newTopSide() *piece {
	if f.theme.DecorSizeTop == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("top", cursors.TopSide)
	PixA, pixI := f.theme.DecorTopA, f.theme.DecorTopI

	win.MROpt(fX|fY|fH, f.theme.DecorSizeTop, 0, 0, f.theme.DecorSizeTop)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newBottomSide() *piece {
	if f.theme.DecorSizeBottom == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("bottom", cursors.BottomSide)
	PixA, pixI := f.theme.DecorBottomA, f.theme.DecorBottomI

	win.MROpt(fX|fH, f.theme.DecorSizeBottom, 0, 0, f.theme.DecorSizeBottom)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newLeftSide() *piece {
	if f.theme.DecorSizeLeft == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("left", cursors.LeftSide)
	PixA, pixI := f.theme.DecorLeftA, f.theme.DecorLeftI

	win.MROpt(fX|fY|fW, 0, f.theme.DecorSizeLeft, f.theme.DecorSizeLeft, 0)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newRightSide() *piece {
	if f.theme.DecorSizeRight == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("right", cursors.RightSide)
	PixA, pixI := f.theme.DecorRightA, f.theme.DecorRightI

	win.MROpt(fY|fW, 0, f.theme.DecorSizeRight, f.theme.DecorSizeRight, 0)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newTopLeft() *piece {
	if f.theme.DecorSizeTop+f.theme.DecorSizeLeft == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("topleft", cursors.TopLeftCorner)
	PixA, pixI := f.theme.DecorTopA, f.theme.DecorTopI

	win.MROpt(fX|fY|fW|fH, 0, 0, f.theme.DecorSizeTop, f.theme.DecorSizeLeft)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newTopRight() *piece {
	if f.theme.DecorSizeTop+f.theme.DecorSizeRight == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("topright", cursors.TopRightCorner)
	PixA, pixI := f.theme.DecorTopA, f.theme.DecorTopI

	win.MROpt(fY|fW|fH, 0, 0, f.theme.DecorSizeTop, f.theme.DecorSizeRight)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newBottomLeft() *piece {
	if f.theme.DecorSizeBottom+f.theme.DecorSizeLeft == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("bottomleft", cursors.BottomLeftCorner)
	PixA, pixI := f.theme.DecorTopA, f.theme.DecorTopI

	win.MROpt(fX|fW|fH, 0, 0, f.theme.DecorSizeBottom, f.theme.DecorSizeLeft)

	return newPiece(win, PixA, pixI)
}

func (f *Decor) newBottomRight() *piece {
	if f.theme.DecorSizeBottom+f.theme.DecorSizeRight == 0 {
		return newEmptyPiece()
	}

	win := f.newPieceWindow("bottomright", cursors.BottomRightCorner)
	PixA, pixI := f.theme.DecorTopA, f.theme.DecorTopI

	win.MROpt(fW|fH, 0, 0, f.theme.DecorSizeBottom, f.theme.DecorSizeRight)

	return newPiece(win, PixA, pixI)
}
