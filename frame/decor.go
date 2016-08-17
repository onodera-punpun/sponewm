package frame

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
)

type Decor struct {
	*frame
	theme *DecorTheme

	topSide, bottomSide, leftSide, rightSide   *piece
	topLeft, topRight, bottomLeft, bottomRight *piece
}

func NewDecor(X *xgbutil.XUtil,
	t *DecorTheme, p *Parent, c Client) (*Decor, error) {

	f, err := newFrame(X, p, c)
	if err != nil {
		return nil, err
	}

	df := &Decor{frame: f, theme: t}

	df.topSide = df.newTopSide()
	df.bottomSide = df.newBottomSide()
	df.leftSide = df.newLeftSide()
	df.rightSide = df.newRightSide()

	df.topLeft = df.newTopLeft()
	df.topRight = df.newTopRight()
	df.bottomLeft = df.newBottomLeft()
	df.bottomRight = df.newBottomRight()

	return df, nil
}

func (f *Decor) Current() bool {
	return f.client.Frame() == f
}

func (f *Decor) Destroy() {
	f.topSide.Destroy()
	f.bottomSide.Destroy()
	f.leftSide.Destroy()
	f.rightSide.Destroy()

	f.topLeft.Destroy()
	f.topRight.Destroy()
	f.bottomLeft.Destroy()
	f.bottomRight.Destroy()

	f.frame.Destroy()
}

func (f *Decor) Off() {
	f.topSide.Unmap()
	f.bottomSide.Unmap()
	f.leftSide.Unmap()
	f.rightSide.Unmap()

	f.topLeft.Unmap()
	f.topRight.Unmap()
	f.bottomLeft.Unmap()
	f.bottomRight.Unmap()
}

func (f *Decor) On() {
	Reset(f)

	if f.client.State() == Active {
		f.Active()
	} else {
		f.Inactive()
	}

	if !f.client.IsMaximized() {
		f.topSide.Map()
		f.bottomSide.Map()
		f.leftSide.Map()
		f.rightSide.Map()

		f.topLeft.Map()
		f.topRight.Map()
		f.bottomLeft.Map()
		f.bottomRight.Map()
	}
}

func (f *Decor) Active() {
	f.State = Active

	f.topSide.Active()
	f.bottomSide.Active()
	f.leftSide.Active()
	f.rightSide.Active()

	f.topLeft.Active()
	f.topRight.Active()
	f.bottomLeft.Active()
	f.bottomRight.Active()

	f.parent.Change(xproto.CwBackPixel, uint32(0xffffff))
	f.parent.ClearAll()
}

func (f *Decor) Inactive() {
	f.State = Inactive

	f.topSide.Inactive()
	f.bottomSide.Inactive()
	f.leftSide.Inactive()
	f.rightSide.Inactive()

	f.topLeft.Inactive()
	f.topRight.Inactive()
	f.bottomLeft.Inactive()
	f.bottomRight.Inactive()

	f.parent.Change(xproto.CwBackPixel, uint32(0xffffff))
	f.parent.ClearAll()
}

func (f *Decor) Maximize() {
	if f.theme.DecorSizeTop+f.theme.DecorSizeBottom+
		f.theme.DecorSizeLeft+
		f.theme.DecorSizeRight > 0 && f.Current() {

		f.topSide.Unmap()
		f.bottomSide.Unmap()
		f.leftSide.Unmap()
		f.rightSide.Unmap()

		f.topLeft.Unmap()
		f.topRight.Unmap()
		f.bottomLeft.Unmap()
		f.bottomRight.Unmap()

		Reset(f)
	}
}

func (f *Decor) Unmaximize() {
	if f.theme.DecorSizeTop+f.theme.DecorSizeBottom+
		f.theme.DecorSizeLeft+f.theme.DecorSizeRight > 0 && f.Current() {

		f.topSide.Map()
		f.bottomSide.Map()
		f.leftSide.Map()
		f.rightSide.Map()

		f.topLeft.Map()
		f.topRight.Map()
		f.bottomLeft.Map()
		f.bottomRight.Map()

		Reset(f)
	}
}

func (f *Decor) Top() int {
	if f.client.IsMaximized() {
		return 0
	}
	return f.theme.DecorSizeTop
}

func (f *Decor) Bottom() int {
	if f.client.IsMaximized() {
		return 0
	}
	return f.theme.DecorSizeBottom
}

func (f *Decor) Left() int {
	if f.client.IsMaximized() {
		return 0
	}
	return f.theme.DecorSizeLeft
}

func (f *Decor) Right() int {
	if f.client.IsMaximized() {
		return 0
	}
	return f.theme.DecorSizeRight
}

func (f *Decor) moveresizePieces() {
	fg := f.Geom()

	f.topSide.MROpt(fW, 0, 0, fg.Width()-f.topLeft.w()-f.topRight.w(), 0)
	f.bottomSide.MROpt(fY|fW, 0, fg.Height()-f.bottomSide.h(), f.topSide.w(), 0)
	f.leftSide.MROpt(fH, 0, 0, 0, fg.Height()-f.topLeft.h()-f.bottomLeft.h())
	f.rightSide.MROpt(fX|fH, fg.Width()-f.rightSide.w(), 0, 0, f.leftSide.h())

	f.topRight.MROpt(fX, f.topLeft.w()+f.topSide.w(), 0, 0, 0)
	f.bottomLeft.MROpt(fY, 0, f.bottomSide.y(), 0, 0)
	f.bottomRight.MROpt(fX|fY,
		f.bottomLeft.w()+f.bottomSide.w(), f.bottomSide.y(), 0, 0)
}

func (f *Decor) MROpt(validate bool, flags, x, y, w, h int) {
	mropt(f, validate, flags, x, y, w, h)
	f.moveresizePieces()
}

func (f *Decor) MoveResize(validate bool, x, y, w, h int) {
	moveresize(f, validate, x, y, w, h)
	f.moveresizePieces()
}

func (f *Decor) Move(x, y int) {
	move(f, x, y)
}

func (f *Decor) Resize(validate bool, w, h int) {
	resize(f, validate, w, h)
	f.moveresizePieces()
}

type DecorTheme struct {
	DecorTopA, DecorTopI       *xgraphics.Image
	DecorBottomA, DecorBottomI *xgraphics.Image
	DecorLeftA, DecorLeftI     *xgraphics.Image
	DecorRightA, DecorRightI   *xgraphics.Image
	DecorSizeTop               int
	DecorSizeBottom            int
	DecorSizeLeft              int
	DecorSizeRight             int
}
