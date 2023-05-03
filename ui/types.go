package ui

import (
	"pixel/apptype"
	"pixel/pxcanvas"
	"pixel/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixelCanvas *pxcanvas.PxCanvas
	PixelWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
