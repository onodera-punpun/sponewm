package wm

import (
	"github.com/onodera-punpun/wingo/frame"
	"github.com/onodera-punpun/wingo/misc"
	"github.com/onodera-punpun/wingo/render"
	"github.com/onodera-punpun/wingo/wini"
)

type ThemeConfig struct {
	Borders     ThemeBorders
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

func newTheme() *ThemeConfig {
	return &ThemeConfig{
		Borders: ThemeBorders{
			borderSize:   10,
			aBorderColor: render.NewColor(0xeeeeee),
			iBorderColor: render.NewColor(0xeeeeee),
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
