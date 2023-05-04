package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func Setup(app *AppInit) {
	SetupMenus(app)
	swatchesContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	brushSelector := NewBrushSelector(app.State, app.PixelCanvas)
	brushSelectorObject := brushSelector.Render()
	brushSelectorObject.Resize(fyne.NewSize(100, 50))

	appLayout := container.NewBorder(nil, swatchesContainer, nil, nil, app.PixelCanvas)

	rightSide := container.NewVBox(colorPicker, brushSelectorObject)
	rightSide.Resize(fyne.NewSize(200, 0))

	horizontalLayout := container.NewHBox(appLayout, rightSide)
	app.PixelWindow.SetContent(horizontalLayout)
}
