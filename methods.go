package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type Results struct { // define a new structure called Results with two fields; result, and pdiff ::: - -
	result float64
	pdiff  float64
}

// Theme
type myTheme struct { // ::: - -
	Theme fyne.Theme
}

func (m *myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{245, 245, 245, 255}
	case theme.ColorNameButton:
		return color.RGBA{200, 200, 200, 255}
	case theme.ColorNameForeground:
		return color.RGBA{0, 0, 0, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 165, 0, 255}
	}
	return m.Theme.Color(name, variant)
}
func (m *myTheme) Font(style fyne.TextStyle) fyne.Resource    { return m.Theme.Font(style) }
func (m *myTheme) Icon(name fyne.ThemeIconName) fyne.Resource { return m.Theme.Icon(name) }
func (m *myTheme) Size(name fyne.ThemeSizeName) float32       { return m.Theme.Size(name) }
