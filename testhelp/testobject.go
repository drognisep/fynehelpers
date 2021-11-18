package testhelp

import (
	"sync"

	"fyne.io/fyne/v2"
)

var _ fyne.CanvasObject = (*testCanvasObject)(nil)

type testCanvasObject struct {
	mux     sync.RWMutex
	visible bool
	minSize fyne.Size
	size    fyne.Size
	pos     fyne.Position
}

func NewTestObject(size fyne.Size) fyne.CanvasObject {
	return &testCanvasObject{
		size:    size,
		minSize: size,
	}
}

func (t *testCanvasObject) MinSize() fyne.Size {
	t.mux.RLock()
	defer t.mux.RUnlock()
	return t.minSize
}

func (t *testCanvasObject) Move(pos fyne.Position) {
	t.mux.Lock()
	t.pos = pos
	t.mux.Unlock()
}

func (t *testCanvasObject) Position() fyne.Position {
	t.mux.RLock()
	defer t.mux.RUnlock()
	return t.pos
}

func (t *testCanvasObject) Resize(size fyne.Size) {
	t.mux.Lock()
	t.size = size
	t.mux.Unlock()
}

func (t *testCanvasObject) Size() fyne.Size {
	t.mux.RLock()
	defer t.mux.RUnlock()
	return t.size
}

func (t *testCanvasObject) Hide() {
	t.mux.Lock()
	t.visible = false
	t.mux.Unlock()
}

func (t *testCanvasObject) Visible() bool {
	t.mux.RLock()
	defer t.mux.RUnlock()
	return t.visible
}

func (t *testCanvasObject) Show() {
	t.mux.Lock()
	t.visible = true
	t.mux.Unlock()
}

func (t *testCanvasObject) Refresh() {
}
