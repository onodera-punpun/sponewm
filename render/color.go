package render

import (
	"fmt"

	"image/color"
)

type Color struct {
	start, end int
}

var NoColor = Color{}

func NewColor(clr int) Color {
	return Color{start: clr, end: -1}
}

func NewImageColor(clr color.Color) Color {
	return NewColor(intFromImageColor(clr))
}

func (c Color) String() string {
	return fmt.Sprintf("%x", c.start)
}

func (c *Color) ColorSet(clr int) {
	c.start = clr
}

func (c Color) Int() int {
	return c.start
}

func (c Color) Uint32() uint32 {
	return uint32(c.start)
}

func (c Color) ImageColor() color.RGBA {
	return imageColorFromInt(c.start)
}

func (c Color) RGB() (r, g, b int) {
	return rgbFromInt(c.start)
}

func (c Color) RGB8() (r, g, b uint8) {
	r32, g32, b32 := c.RGB()
	r, g, b = uint8(r32), uint8(g32), uint8(b32)
	return
}

func intFromImageColor(clr color.Color) int {
	r, g, b, _ := clr.RGBA()
	return intFromRgb(int(r>>8), int(g>>8), int(b>>8))
}

func imageColorFromInt(clr int) color.RGBA {
	r, g, b := rgbFromInt(clr)
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func intFromRgb(r, g, b int) int {
	return (r << 16) + (g << 8) + b
}

func rgbFromInt(clr int) (r, g, b int) {
	r = clr >> 16
	g = (clr & 0x00ff00) >> 8
	b = clr & 0x0000ff
	return
}
