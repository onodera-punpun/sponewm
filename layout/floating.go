package layout

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xrect"
)

type Floating struct {
	clients []Client
	geom    xrect.Rect
}

func NewFloating() *Floating {
	return &Floating{
		clients: make([]Client, 0),
	}
}

// TODO: Add "border" to monitor edges.
func (f *Floating) InitialPlacement(x *xgbutil.XUtil, c Client) {
	cgeom := c.Geom()
	xc, rw := x.Conn(), x.RootWin()
	qp, _ := xproto.QueryPointer(xc, rw).Reply()

	fx, fy := f.geom.X(), f.geom.Y()
	// TODO: Is this x/y-limit needed, what does it do?
	xlimit := f.geom.Width() - cgeom.Width()
	ylimit := f.geom.Height() - cgeom.Height()
	if xlimit > 0 {
		// TODO: Currently hardcoding the bordersize (10), because I can't
		// seem to get the value
		fx += int(qp.RootX) - cgeom.Width() / 2 + 10
	}
	if ylimit > 0 {
		// TODO: Currently hardcoding the bordersize (10), because I can't
		// seem to get the value
		fy += int(qp.RootY) - cgeom.Height() / 2 + 10
	}
	f.Move(c, fx, fy)
}

func (f *Floating) Place()   {}
func (f *Floating) Unplace() {}
func (f *Floating) Destroy() {}

func (f *Floating) Name() string {
	return "Floating"
}

func (f *Floating) SetGeom(geom xrect.Rect) {
	f.geom = geom
}

// Save is called when a workspace switches from a floating layout to a
// tiling layout. It should save the "last-floating" state for all floating
// clients.
func (f *Floating) Save() {
	for _, c := range f.clients {
		if _, ok := c.Layout().(*Floating); ok {
			c.SaveState("last-floating")
		}
	}
}

// Reposition is called when a workspace switches from a tiling layout to a
// floating layout. It should reload the "last-floating" client state.
func (f *Floating) Reposition() {
	if f.geom == nil {
		return
	}
	for _, c := range f.clients {
		// Don't reposition windows that are already in the floating layout.
		if c.ShouldForceFloating() {
			continue
		}
		if _, ok := c.Layout().(*Floating); ok {
			c.LoadState("last-floating")
		}
	}
}

func (f *Floating) Exists(c Client) bool {
	for _, client := range f.clients {
		if client == c {
			return true
		}
	}
	return false
}

func (f *Floating) Add(c Client) {
	if !f.Exists(c) {
		f.clients = append(f.clients, c)
	}
}

func (f *Floating) Remove(c Client) {
	for i, client := range f.clients {
		if client == c {
			f.clients = append(f.clients[:i], f.clients[i+1:]...)
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
