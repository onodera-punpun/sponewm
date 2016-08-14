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

func (t *Tiling) Place(padding []int, gap int) {
	// TODO: Is the 0 needed here?
	x := 0 + padding[0]
	y := 0 + padding[2]
	width := t.geom.Width() + t.geom.X() - padding[1] - padding[3]
	height := t.geom.Height() + t.geom.Y() - padding[0] - padding[2]

	n := t.clients.Len() - 1
	i := 0

	for l := t.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)

		if i < n {
			if i % 2 != 0 {
				height /= 2
			} else {
				width /= 2
			}
			if i % 4 == 2 {
				x += width
			} else if i % 4 == 3 {
				y += height
			}
		}
		if i % 4 == 0 {
			y -= height
		} else if i % 4 == 1 {
			x += width
		} else if i % 4 == 2 {
			y += height
		} else if i % 4 == 3 {
			x -= width
		}
		if i == 0 {
			if n != 0 {
				width = int(float64(t.geom.Width() + t.geom.X()) * 0.617) - padding[1] - padding[3]
			}
			y = 0 + padding[2]
		} else if i == 1 {
			width = t.geom.Width() + t.geom.X() - padding[1] - padding[3] - width
		}

		c.FrameTile()
		c.MoveResize(x + gap, y + gap, width - gap, height - gap)
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

func (t *Tiling) MakeMaster(c Client) {
	t.clients.PushFront(c)
}
