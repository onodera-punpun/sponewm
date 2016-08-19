package wm

type CommandHacks struct {
	MouseResizeDirection     func(cmdStr string) (string, error)
}
