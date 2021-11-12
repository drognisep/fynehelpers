package layouthelp

import "fyne.io/fyne/v2"

func AccumulateWidth(sizes ...fyne.Size) fyne.Size {
	var width float32 = 0.0
	var maxHeight float32 = 0.0
	for _, sz := range sizes {
		if sz.Height > maxHeight {
			maxHeight = sz.Height
		}
		width += sz.Width
	}
	return fyne.NewSize(width, maxHeight)
}

func AccumulateHeight(sizes ...fyne.Size) fyne.Size {
	var height float32 = 0.0
	var maxWidth float32 = 0.0
	for _, sz := range sizes {
		if sz.Width > maxWidth {
			maxWidth = sz.Width
		}
		height += sz.Height
	}
	return fyne.NewSize(maxWidth, height)
}
