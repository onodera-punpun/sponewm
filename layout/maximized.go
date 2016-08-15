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
	for l := m.clients.Front(); l != nil; l = l.Next() {
		c := l.Value.(Client)

		c.FrameTile()
		c.MoveResize(m.geom.X(), m.geom.Y(), m.geom.Width(), m.geom.Height())
	}
}

func (m *Maximized) Unplace() {}

func (m *Maximized) Add(c Client) {
	if !m.Exists(c) {
		m.clients.PushFront(c)
	}
}

func (t *Maximized) Remove(c Client) {}

func (m *Maximized) Exists(c Client) bool {
	for l := m.clients.Front(); l != nil; l = l.Next() {
		if l.Value.(Client) == c {
			return true
		}
	}
	return false
}

func (m *Maximized) Destroy() {}

func (m *Maximized) Save() {}

func (m *Maximized) Reposition() {}

func (m *Maximized) MROpt(c Client, flags, x, y, width, height int) {}

func (m *Maximized) MoveResize(c Client, x, y, width, height int) {}

func (m *Maximized) Move(c Client, x, y int) {}

func (m *Maximized) Resize(c Client, width, height int) {}

func (m *Maximized) MakeMaster() {}