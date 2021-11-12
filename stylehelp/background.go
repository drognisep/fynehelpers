package stylehelp

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func AddBackgroundColor(color color.NRGBA, obj fyne.CanvasObject) fyne.CanvasObject {
	rect := canvas.NewRectangle(color)
	return container.NewMax(rect, obj)
}
