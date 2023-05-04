package brush

import (
	"image/color"
	"pixel/apptype"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
)

const (
	Pixel = iota
	Fat
)

func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, ev *desktop.MouseEvent, x int, y int) []fyne.CanvasObject {
	var objects []fyne.CanvasObject
	switch {
	case brush == Pixel:
		pxSize := float32(config.PxSize)
		xOrigin := (float32(x) * pxSize) + config.CanvasOffset.X
		yOrigin := (float32(y) * pxSize) + config.CanvasOffset.Y

		cursorColor := color.NRGBA{80, 80, 80, 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize)
		bottom.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		objects = append(objects, left, top, right, bottom)
	case brush == Fat:
		pxSize := float32(config.PxSize) * 2 // set pixel size to twice the value for Fat brush
		xOrigin := (float32(x) * pxSize / 2) + config.CanvasOffset.X - pxSize/2
		yOrigin := (float32(y) * pxSize / 2) + config.CanvasOffset.Y - pxSize/2

		cursorColor := color.NRGBA{80, 80, 80, 255}

		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize)
		bottom.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		objects = append(objects, left, top, right, bottom)
	}
	return objects
}

func TryBrush(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	case appState.BrushType == Fat: // added case statement for Fat brush
		return TryPaintFat(appState, canvas, ev)
	default:
		return false
	}
}

func TryPaintPixel(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}

func TryPaintFat(appState *apptype.State, canvas apptype.Brushable, ev *desktop.MouseEvent) bool {
	x, y := canvas.MouseToCanvasXY(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		// Set the color of the pixels in a 2x2 square centered on the cursor position
		canvas.SetColor(appState.BrushColor, *x, *y)
		canvas.SetColor(appState.BrushColor, *x-1, *y)
		canvas.SetColor(appState.BrushColor, *x, *y-1)
		canvas.SetColor(appState.BrushColor, *x-1, *y-1)
		return true
	}
	return false
}
