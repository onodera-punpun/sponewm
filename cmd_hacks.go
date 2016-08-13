package main

import (
	"github.com/onodera-punpun/sponewm/commands"
	"github.com/onodera-punpun/sponewm/wm"
)

func newHacks() wm.CommandHacks {
	return wm.CommandHacks{
		MouseResizeDirection:     mouseResizeDirection,
	}
}

func mouseResizeDirection(cmdStr string) (string, error) {
	cmd, err := commands.Env.Command(cmdStr)
	if err != nil {
		return "", err
	}
	return cmd.(*commands.MouseResize).Direction, nil
}
