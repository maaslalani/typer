package theme

import (
	"github.com/charmbracelet/bubbles/progress"
	"github.com/guptarohit/asciigraph"
	"github.com/muesli/termenv"
	"github.com/spf13/viper"
)

func DefaultTheme() *Theme {
	return &Theme{
		Text: Text{
			Untyped: RGBColor{
				Foreground: "#555",
				Background: "",
			},
			Typed: RGBColor{
				Foreground: "#fff",
				Background: "",
			},
			Error: RGBColor{
				Foreground: "#fff",
				Background: "#f33",
			},
		},
		Bar: Bar{
			Color:    "#4776E6",
			Gradient: "",
		},
		Graph: Graph{
			Color:  "blue",
			Height: 3,
		},
	}
}

func LoadViper(v *viper.Viper, first bool) (*Theme, error) {
	theme := DefaultTheme()
	v.UnmarshalKey("theme", theme)
	if !first || theme.File == "" {
		return theme, nil
	}

	v = viper.New()
	v.SetConfigFile(theme.File)
	err := v.ReadInConfig()
	if err != nil {
		return theme, err
	}
	return LoadViper(v, false)
}

type Theme struct {
	File  string `json:"file" toml:"file" yaml:"file"`
	Text  Text   `json:"text" toml:"text" yaml:"text"`
	Bar   Bar    `json:"bar" toml:"bar" yaml:"bar"`
	Graph Graph  `json:"graph" toml:"graph" yaml:"graph"`
}

type Text struct {
	Untyped         RGBColor `json:"untyped" toml:"untyped" yaml:"untyped"`
	Typed           RGBColor `json:"typed" toml:"typed" yaml:"typed"`
	Error           RGBColor `json:"error" toml:"error" yaml:"error"`
	ErrorForeground RGBColor `json:"error_foreground" toml:"error_foreground" yaml:"error_foreground"`
}

type RGBColor struct {
	Foreground string `json:"foreground" toml:"foreground" yaml:"foreground"`
	Background string `json:"background" toml:"background" yaml:"background"`
}

type Bar struct {
	Color    string `json:"color" toml:"color" yaml:"color"`
	Gradient string `json:"gradient" toml:"gradient" yaml:"gradient"`
}

type Graph struct {
	Color  string `json:"color" toml:"color" yaml:"color"`
	Height int    `json:"height" toml:"height" yaml:"height"`
}

func (t Theme) StringColor(rgbColors RGBColor, input string) termenv.Style {
	output := termenv.String(input)

	if rgbColors.Background != "" {
		output = output.Background(termenv.RGBColor(rgbColors.Background))
	}

	if rgbColors.Foreground != "" {
		output = output.Foreground(termenv.RGBColor(rgbColors.Foreground))
	}

	return output
}

func (t Theme) BarColor() progress.Option {
	if t.Bar.Color == "" {
		return progress.WithDefaultGradient()
	}

	if t.Bar.Gradient == "" {
		return progress.WithSolidFill(t.Bar.Color)
	}

	return progress.WithGradient(t.Bar.Color, t.Bar.Gradient)
}

func (t Theme) GraphColor() asciigraph.AnsiColor {
	graphColor := asciigraph.Blue
	if color, ok := asciigraph.ColorNames[t.Graph.Color]; ok {
		graphColor = color
	}
	return graphColor
}
