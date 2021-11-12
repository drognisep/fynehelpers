package layouthelp

import "fyne.io/fyne/v2"

// AccumulateWidth will add all widths together and return a fyne.Size with the sum width and max height.
// This is useful for defining MinSize of fyne.CanvasObjects in a row.
func AccumulateWidth(sizes ...fyne.Size) fyne.Size {
	var width float32
	var maxHeight float32
	for _, sz := range sizes {
		if sz.Height > maxHeight {
			maxHeight = sz.Height
		}
		width += sz.Width
	}
	return fyne.NewSize(width, maxHeight)
}

// AccumulateHeight will add all heights together and return a fyne.Size with the sum height and max width.
// This is useful for defining MinSize of fyne.CanvasObjects in a column.
func AccumulateHeight(sizes ...fyne.Size) fyne.Size {
	var height float32
	var maxWidth float32
	for _, sz := range sizes {
		if sz.Width > maxWidth {
			maxWidth = sz.Width
		}
		height += sz.Height
	}
	return fyne.NewSize(maxWidth, height)
}
