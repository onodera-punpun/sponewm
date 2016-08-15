package layout

import (
	"container/list"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xrect"
)

type Floating struct {
	clients *list.List
	geom    xrect.Rect
}

func NewFloating() *Floating {
	return &Floating{
		clients: list.New(),
	}
}

func (f *Floating) Name() string {
	return "Floating"
}

func (f *Floating) SetGeom(geom xrect.Rect) {
	f.geom = geom
}

func (f *Floating) Place()   {}

func (f *Floating) Unplace() {}

func (f *Floating) Add(c Client) {
	if !f.Exists(c) {
		f.clients.PushFront(c)
	}
}

func (f *Floating) Remove(c Client) {
	for l := f.clients.Front(); l != nil; l = l.Next() {
		if l.Value.(Client) == c {
			f.clients.Remove(l)
		}
	}
}

func (f *Floating) Exists(c Client) bool {
	for l := f.clients.Front(); l != nil; l = l.Next() {
		if l.Value.(Client) == c {
			return true
		}
	}
	return false
}

func (f *Floating) Destroy() {}

// Save is called when a workspace switches from a floating layout to a
// tiling layout. It should save the "last-floating" state for all floating
// clients.
func (f *Floating) Save() {
	for l := f.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)
		if _, ok := c.Layout().(*Floating); ok {
			c.SaveState("last-floating")
		}
	}
}

// Reposition is called when a workspace switches from a tiling layout to a
// floating layout. It should reload the "last-floating" client state.
func (f *Floating) Reposition() {
	for l := f.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)
		if _, ok := c.Layout().(*Floating); ok {
			c.LoadState("last-floating")
		}
	}
}

func (f *Floating) MROpt(c Client, flags, x, y, width, height int) {
	c.MROpt(true, flags, x, y, width, height)
	c.SaveState("last-floating")
}

func (f *Floating) MoveResize(c Client, x, y, width, height int) {
	c.MoveResizeValid(x, y, width, height)
	c.SaveState("last-floating")
}

func (f *Floating) Move(c Client, x, y int) {
	c.Move(x, y)
	c.SaveState("last-floating")
}

func (f *Floating) Resize(c Client, width, height int) {
	c.Resize(true, width, height)
	c.SaveState("last-floating")
}

func (f *Floating) InitialPlacement(c Client, X *xgbutil.XUtil) {
	// TODO: I'm gonna hardcode this because I can't
	// figure out this circular depenency shit.
	padding := []int{29, 20, 20, 20}

	cgeom := c.Geom()
	qp, _ := xproto.QueryPointer(X.Conn(), X.RootWin()).Reply()

	x := int(qp.RootX) - (cgeom.Width() / 2)
	y := int(qp.RootY) - (cgeom.Height() / 2)
	// Left screen border.
	if x < padding[3] {
		x = padding[3]
	// Right screen border.
	} else if x > f.geom.Width() + f.geom.X() - cgeom.Width() - padding[1] {
		x = f.geom.Width() + f.geom.X() - cgeom.Width() - padding[1]
	}
	// Top screen border.
	if y < padding[0] {
		y = padding[0]
	// Bottom screen border.
	} else if y > f.geom.Height() + f.geom.Y() - cgeom.Height() - padding[2] {
		y = f.geom.Height() + f.geom.Y() - cgeom.Height() - padding[2] 
	}

	f.Move(c, x, y)
}
