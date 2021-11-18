package layouthelp

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

// MultiView allows different fyne.CanvasObjects to be displayed in the same area of a layout. It can be used as either
// a stack which allows pushing an object for a separate task and popping it off to return to the previous view, or
// as a replacement target that discards the existing canvas object instead.
type MultiView struct {
	container fyne.Container
	mux       sync.Mutex
	viewStack []fyne.CanvasObject
}

// NewMultiView should be used to initialize a MultiView.
func NewMultiView() *MultiView {
	mv := &MultiView{}
	mv.container.Layout = layout.NewMaxLayout()
	return mv
}

// Push adds the provided canvas object to the top of the view stack and replaces the container content with it.
func (v *MultiView) Push(obj fyne.CanvasObject) {
	defer v.container.Refresh()
	v.mux.Lock()
	defer v.mux.Unlock()
	v.push(obj)
}

func (v *MultiView) push(obj fyne.CanvasObject) {
	stackDepth := len(v.viewStack)
	if stackDepth > 0 {
		v.container.Remove(v.viewStack[stackDepth-1])
	}
	v.viewStack = append(v.viewStack, obj)
	v.container.Add(obj)
}

// Pop pops the top canvas object off the view stack and replaces the container content with the next if there is one.
func (v *MultiView) Pop() {
	defer v.container.Refresh()
	v.mux.Lock()
	defer v.mux.Unlock()
	v.pop()
}

func (v *MultiView) pop() {
	stackDepth := len(v.viewStack)
	if stackDepth == 0 {
		return
	}
	v.container.Remove(v.viewStack[stackDepth-1])
	v.viewStack = v.viewStack[0 : stackDepth-1]
	if len(v.viewStack) > 0 {
		v.container.Add(v.viewStack[stackDepth-2])
	}
}

// Container gets the configured container which can be manipulated through the MultiView.
func (v *MultiView) Container() *fyne.Container {
	return &v.container
}

// Clear removes all objects from the stack and the container.
func (v *MultiView) Clear() {
	v.mux.Lock()
	defer v.mux.Unlock()
	last := v.viewStack[len(v.viewStack)-1]
	v.viewStack = nil
	v.container.Remove(last)
}

// Replace replaces the currently displayed fyne.CanvasObject (if any) with the provided object.
func (v *MultiView) Replace(obj fyne.CanvasObject) {
	defer v.container.Refresh()
	v.mux.Lock()
	defer v.mux.Unlock()
	v.pop()
	v.push(obj)
}
