package layout

import (
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xrect"
)

type Layout interface {
	Name() string
	SetGeom(geom xrect.Rect)

	Place()
	Unplace()

	Add(c Client)
	Remove(c Client)
	Exists(c Client) bool

	Destroy()

	Save()
	Reposition()

	MROpt(c Client, flags, x, y, width, height int)
	MoveResize(c Client, x, y, width, height int)
	Move(c Client, x, y int)
	Resize(c Client, width, height int)
}

type Floater interface {
	Layout
	InitialPlacement(c Client, X *xgbutil.XUtil)
}

type Tiler interface {
	Layout
	MakeMaster()
}
