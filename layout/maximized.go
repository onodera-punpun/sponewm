package layout

import (
	"container/list"

	"github.com/BurntSushi/xgbutil/xrect"
)

type Maximized struct {
	clients *list.List
	geom    xrect.Rect
}

func NewMaximized() *Maximized {
	return &Maximized{
		clients: list.New(),
	}
}

func (m *Maximized) Name() string {
	return "Maximized"
}

func (m *Maximized) SetGeom(geom xrect.Rect) {
	m.geom = geom
}

func (m *Maximized) Place() {
	for el := m.clients.Front(); el != nil; el = el.Next() {
		c := el.Value.(Client)
		x, y, w, h := m.geom.X(), m.geom.Y(), m.geom.Width(), m.geom.Height()
		c.FrameTile()
		c.MoveResize(x, y, w, h)
	}
}

func (m *Maximized) Add(c Client) {
	if !m.Exists(c) {
		m.clients.PushFront(c)
	}
}

func (m *Maximized) Exists(c Client) bool {
	for e := m.clients.Front(); e != nil; e = e.Next() {
		if e.Value.(Client) == c {
			return true
		}
	}
	return false
}

func (m *Maximized) MROpt(c Client, flags, x, y, width, height int) {}

func (m *Maximized) MoveResize(c Client, x, y, width, height int) {}

func (m *Maximized) Move(c Client, x, y int) {}

func (m *Maximized) Resize(c Client, width, height int) {}
