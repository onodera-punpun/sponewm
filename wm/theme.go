package wm

import (
	"github.com/onodera-punpun/wingo/frame"
	"github.com/onodera-punpun/wingo/misc"
	"github.com/onodera-punpun/wingo/render"
	"github.com/onodera-punpun/wingo/wini"
)

type ThemeConfig struct {
	Borders     ThemeBorders
	Slim        ThemeSlim
}

type ThemeBorders struct {
	borderSize                 int
	aBorderColor, iBorderColor render.Color
}

func (tb ThemeBorders) FrameTheme() *frame.BordersTheme {
	return &frame.BordersTheme{
		BorderSize:   tb.borderSize,
		ABorderColor: tb.aBorderColor,
		IBorderColor: tb.iBorderColor,
	}
}

type ThemeSlim struct {
	borderSize                 int
	aBorderColor, iBorderColor render.Color
}

func (ts ThemeSlim) FrameTheme() *frame.SlimTheme {
	return &frame.SlimTheme{
		BorderSize:   ts.borderSize,
		ABorderColor: ts.aBorderColor,
		IBorderColor: ts.iBorderColor,
	}
}

func newTheme() *ThemeConfig {
	return &ThemeConfig{
		Borders: ThemeBorders{
			borderSize:   10,
			aBorderColor: render.NewColor(0xeeeeee),
			iBorderColor: render.NewColor(0xeeeeee),
		},
		Slim: ThemeSlim{
			borderSize:   10,
			aBorderColor: render.NewColor(0xeeeee),
			iBorderColor: render.NewColor(0xeeeee),
		},
	}
}

func loadTheme() (*ThemeConfig, error) {
	theme := newTheme()

	tdata, err := wini.Parse(misc.ConfigFile("theme.wini"))
	if err != nil {
		return nil, err
	}

	for _, section := range tdata.Sections() {
		switch section {
		case "borders":
			for _, key := range tdata.Keys(section) {
				loadBorderOption(theme, key)
			}
		case "slim":
			for _, key := range tdata.Keys(section) {
				loadSlimOption(theme, key)
			}
		}
	}

	return theme, nil
}

func loadBorderOption(theme *ThemeConfig, k wini.Key) {
	switch k.Name() {
	case "border_size":
		setInt(k, &theme.Borders.borderSize)
	case "a_border_color":
		setColor(k, &theme.Borders.aBorderColor)
	case "i_border_color":
		setColor(k, &theme.Borders.iBorderColor)
	}
}

func loadSlimOption(theme *ThemeConfig, k wini.Key) {
	switch k.Name() {
	case "border_size":
		setInt(k, &theme.Slim.borderSize)
	case "a_border_color":
		setColor(k, &theme.Slim.aBorderColor)
	case "i_border_color":
		setColor(k, &theme.Slim.iBorderColor)
	}
}
