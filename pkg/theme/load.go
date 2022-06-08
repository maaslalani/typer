package theme

import (
	"github.com/spf13/viper"
)

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
