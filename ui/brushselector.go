package ui

import (
	"pixel/apptype"
	"pixel/pxcanvas/brush"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type BrushSelector struct {
	appState *apptype.State
	canvas   apptype.Brushable
}

func NewBrushSelector(appState *apptype.State, canvas apptype.Brushable) *BrushSelector {
	return &BrushSelector{appState: appState, canvas: canvas}
}

func (s *BrushSelector) PixelButtonTapped() {
	s.appState.BrushType = brush.Pixel
}

func (s *BrushSelector) FatButtonTapped() {
	s.appState.BrushType = brush.Fat
}

func (s *BrushSelector) Render() fyne.CanvasObject {
	pixelButton := widget.NewButton("Pixel Brush", s.PixelButtonTapped)
	fatButton := widget.NewButton("Fat Brush", s.FatButtonTapped)

	buttons := container.NewGridWithColumns(2, pixelButton, fatButton)
	return container.NewHBox(buttons)
}
