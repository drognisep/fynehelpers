package layouthelp

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var _ fyne.Layout = (*marginLayout)(nil)

type marginLayout struct {
	left, right, top, bottom float32
}

func (p *marginLayout) Layout(objs []fyne.CanvasObject, _ fyne.Size) {
	if len(objs) == 0 {
		return
	}
	objs[0].Move(fyne.NewPos(p.left, p.top))
}

func (p *marginLayout) MinSize(objs []fyne.CanvasObject) fyne.Size {
	if len(objs) == 0 {
		return fyne.Size{}
	}
	objSz := objs[0].MinSize()
	return fyne.NewSize(p.left+p.right+objSz.Width, p.top+p.bottom+objSz.Height)
}

// MarginAll adds the specified padding on all sides of the object.
func MarginAll(obj fyne.CanvasObject, all float32) fyne.CanvasObject {
	lyt := &marginLayout{
		top:    all,
		bottom: all,
		left:   all,
		right:  all,
	}
	return container.New(lyt, obj)
}

// MarginXY adds padding to both sides of each axis, 'x' specifies horizontal padding and 'y' specifies vertical.
func MarginXY(obj fyne.CanvasObject, x, y float32) fyne.CanvasObject {
	lyt := &marginLayout{
		top:    y,
		bottom: y,
		left:   x,
		right:  x,
	}
	return container.New(lyt, obj)
}

// MarginLeftRight allows specifying left and right padding, but leaves top and bottom at zero.
func MarginLeftRight(obj fyne.CanvasObject, left, right float32) fyne.CanvasObject {
	lyt := &marginLayout{
		left:  left,
		right: right,
	}
	return container.New(lyt, obj)
}

// MarginTopBottom allows specifying top and bottom padding, but leaves left and right at zero.
func MarginTopBottom(obj fyne.CanvasObject, top, bottom float32) fyne.CanvasObject {
	lyt := &marginLayout{
		top:    top,
		bottom: bottom,
	}
	return container.New(lyt, obj)
}

// Margin allows specifying padding individually on all sides.
func Margin(obj fyne.CanvasObject, top, right, bottom, left float32) fyne.CanvasObject {
	lyt := &marginLayout{
		top:    top,
		bottom: bottom,
		left:   left,
		right:  right,
	}
	return container.New(lyt, obj)
}
