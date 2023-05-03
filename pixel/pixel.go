package main

import (
	"image/color"
	"pixel/apptype"
	"pixel/pxcanvas"
	"pixel/swatch"
	"pixel/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	pixelApp := app.New()
	pixelWindow := pixelApp.NewWindow("pixel")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	pixelCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	pixelCanvas := pxcanvas.NewPxCanvas(&state, pixelCanvasConfig)

	appInit := ui.AppInit{
		PixelCanvas: pixelCanvas,
		PixelWindow: pixelWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixelWindow.ShowAndRun()
}
