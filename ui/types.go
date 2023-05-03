package ui

import (
	"pixel/apptype"
	"pixel/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
