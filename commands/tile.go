package commands

import (
	"github.com/BurntSushi/gribble"

	"github.com/onodera-punpun/sponewm/workspace"
	"github.com/onodera-punpun/sponewm/xclient"
)

type Tile struct {
	Workspace gribble.Any `param:"1" types:"int,string"`
	Help      string      `
Initiates tiling on the workspace specified by Workspace. If tiling
is already active, the layout will be re-placed.

Note that this command has no effect if the workspace is not visible.

Workspace may be a workspace index (integer) starting at 0, or a workspace
name.
`
}

func (cmd Tile) Run() gribble.Value {
	return syncRun(func() gribble.Value {
		withWorkspace(cmd.Workspace, func(wrk *workspace.Workspace) {
			wrk.LayoutStateSet(workspace.Tiling)
		})
		return nil
	})
}

type Untile struct {
	Workspace gribble.Any `param:"1" types:"int,string"`
	Help      string      `
Stops tiling on the workspace specified by Workspace, and restores
windows to their position and geometry before being tiled. If tiling is not
active on the specified workspace, this command has no effect.

Note that this command has no effect if the workspace is not visible.

Workspace may be a workspace index (integer) starting at 0, or a workspace
name.
`
}

func (cmd Untile) Run() gribble.Value {
	return syncRun(func() gribble.Value {
		withWorkspace(cmd.Workspace, func(wrk *workspace.Workspace) {
			wrk.LayoutStateSet(workspace.Floating)
		})
		return nil
	})
}

type TileToggle struct {
	Workspace gribble.Any `param:"1" types:"int,string"`
	Help      string      `
Toggles tiling, see Tile and Untile.
`
}

func (cmd TileToggle) Run() gribble.Value {
	return syncRun(func() gribble.Value {
		withWorkspace(cmd.Workspace, func(wrk *workspace.Workspace) {
			if wrk.State != workspace.Tiling {
				wrk.LayoutStateSet(workspace.Tiling)
			} else {
				wrk.LayoutStateSet(workspace.Floating)
			}
		})
		return nil
	})
}

type MakeMaster struct {
	Workspace gribble.Any `param:"1" types:"int,string"`
	Client gribble.Any `param:"2" types:"int,string"`
	Help      string      `
Switches the current window with the first master in the layout for the
workspace specified by Workspace.

Note that this command has no effect if the workspace is not visible.

Workspace may be a workspace index (integer) starting at 0, or a workspace
name.
`
}

func (cmd MakeMaster) Run() gribble.Value {
	return syncRun(func() gribble.Value {
		withWorkspace(cmd.Workspace, func(wrk *workspace.Workspace) {
			if wrk.State != workspace.Tiling {
				return
			}
			withClient(cmd.Client, func(c *xclient.Client) {
				wrk.LayoutTiler().MakeMaster(c)
			})
		})
		return nil
	})
}
