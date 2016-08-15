package layout

import (
	"container/list"

	"github.com/BurntSushi/xgbutil/xrect"
)

type Tiling struct {
	clients *list.List
	geom    xrect.Rect
}

func NewTiling() *Tiling {
	return &Tiling{
		clients: list.New(),
	}
}

func (t *Tiling) Name() string {
	return "Tiling"
}

func (t *Tiling) SetGeom(geom xrect.Rect) {
	t.geom = geom
}

func (t *Tiling) Place() {
	// TODO: I'm gonna hardcode this because I can't
	// figure out this circular depenency shit.
	gap := 20
	padding := 40 - (gap / 2)

	x := t.geom.X() + padding
	y := t.geom.Y() + padding
	width := t.geom.Width() - (padding * 2)
	height := t.geom.Height() - (padding * 2)

	n := t.clients.Len() - 1
	i := 0

	for l := t.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)

		if i < n {
			if i%2 != 0 {
				height /= 2
			} else {
				width /= 2
			}
			if i%4 == 2 {
				x += width
			} else if i%4 == 3 {
				y += height
			}
		}
		if i%4 == 0 {
			y -= height
		} else if i%4 == 1 {
			x += width
		} else if i%4 == 2 {
			y += height
		} else if i%4 == 3 {
			x -= width
		}
		if i == 0 {
			if n != 0 {
				width = int(float64(t.geom.Width())*0.617) - (padding * 2)
			}
			y = t.geom.Y() + padding
		} else if i == 1 {
			width = t.geom.Width() - width - (padding * 2)
		}

		c.FrameTile()
		c.MoveResize(x+(gap/2), y+(gap/2), width-gap, height-gap)
		i++
	}
}

func (t *Tiling) Unplace() {}

func (t *Tiling) Add(c Client) {
	if !t.Exists(c) {
		t.clients.PushFront(c)
	}
}

func (t *Tiling) Remove(c Client) {
	for l := t.clients.Front(); l != nil; l = l.Next() {
		if l.Value.(Client) == c {
			t.clients.Remove(l)
		}
	}
}

func (t *Tiling) Exists(c Client) bool {
	for l := t.clients.Front(); l != nil; l = l.Next() {
		if l.Value.(Client) == c {
			return true
		}
	}
	return false
}

func (t *Tiling) Destroy() {}

// Save is called when a workspace switches from a tiling layout to a
// floating layout. It should save the "last-tiling" state for all tiling
// clients.
func (t *Tiling) Save() {
	for l := t.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)
		if _, ok := c.Layout().(*Tiling); ok {
			c.SaveState("last-tiling")
		}
	}
}

// Reposition is called when a workspace switches from a floating layout to a
// tiling layout. It should reload the "last-tiling" client state.
func (t *Tiling) Reposition() {
	for l := t.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)
		if _, ok := c.Layout().(*Tiling); ok {
			c.LoadState("last-tiling")
		}
	}
}

func (t *Tiling) MROpt(c Client, flags, x, y, width, height int) {}

func (t *Tiling) MoveResize(c Client, x, y, width, height int) {}

func (t *Tiling) Move(c Client, x, y int) {}

func (t *Tiling) Resize(c Client, width, height int) {}

func (t *Tiling) MakeMaster() {
	// TODO
}
