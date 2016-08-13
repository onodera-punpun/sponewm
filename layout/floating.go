package layout

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xrect"

	"github.com/onodera-punpun/wingo/logger"
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

func (f *Floating) InitialPlacement(X *xgbutil.XUtil, c Client, padding []int) {
	cgeom := c.Geom()
	qp, err := xproto.QueryPointer(X.Conn(), X.RootWin()).Reply()
	if err != nil {
		logger.Warning.Printf("Could not query pointer: %s", err)
		return
	}

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
