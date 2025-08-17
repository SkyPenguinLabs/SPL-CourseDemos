//////////////////////////////////////////////////////////// README //////////////////////////////////////////////////////////////////
/*			/////////////////
//////////////
			_______ _     _ __   __  _____  _______ __   _  ______ _     _ _____ __   _        _______ ______  _______
			|______ |____/    \_/   |_____] |______ | \  | |  ____ |     |   |   | \  | |      |_____| |_____] |______
			______| |    \_    |    |       |______ |  \_| |_____| |_____| __|__ |  \_| |_____ |     | |_____] ______|
//////////
  [ + ] |
  [ + ] |  This is designed for people to read! :D
  [ + ] |
  [ + ] |
  [ + ] | 				> This is just a theme manager for implementing a standard dark theme on the GUI that follows a similar design
  [ + ] |				  and color scheme as our other applications and websites do. Not much to cover here, just norm fyne dev.
  [ + ] |
/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/
package Core

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DarkTheme struct{}

func (m DarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 24, G: 24, B: 27, A: 255}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 226, G: 232, B: 240, A: 255}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 39, G: 39, B: 42, A: 255}
	case theme.ColorNameButton:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 255}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 99, G: 102, B: 241, A: 255}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0, G: 0, B: 0, A: 64}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 100, G: 100, B: 100, A: 255}
	default:
		return theme.DefaultTheme().Color(name, variant)
	}
}

func (m DarkTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	case theme.SizeNamePadding:
		return 8
	case theme.SizeNameInnerPadding:
		return 12
	default:
		return theme.DefaultTheme().Size(name)
	}
}

func (m DarkTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m DarkTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}
