package layouthelp

import (
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"github.com/drognisep/fynehelpers/testhelp"
	testify "github.com/stretchr/testify/require"
)

func TestMultiView_Push(t *testing.T) {
	assert := testify.New(t)
	mv := NewMultiView()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)

	a := test.NewApp()
	w := a.NewWindow("")
	w.SetContent(mv.Container())
	w.ShowAndRun()

	lbl := widget.NewLabel("")
	mv.Push(lbl)
	assert.Len(mv.viewStack, 1)
	assert.Len(mv.container.Objects, 1)

	entry := widget.NewEntry()
	mv.Push(entry)
	assert.Len(mv.viewStack, 2)
	assert.Len(mv.container.Objects, 1)
}

func TestMultiView_Pop(t *testing.T) {
	assert := testify.New(t)
	mv := NewMultiView()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)

	a := test.NewApp()
	w := a.NewWindow("")
	w.SetContent(mv.Container())
	w.ShowAndRun()

	lbl := widget.NewLabel("")
	mv.Push(lbl)

	entry := widget.NewEntry()
	mv.Push(entry)

	mv.Pop()
	assert.Len(mv.viewStack, 1)
	assert.Len(mv.container.Objects, 1)

	mv.Pop()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)
}

func TestMultiView_Clear(t *testing.T) {
	assert := testify.New(t)
	mv := NewMultiView()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)

	a := test.NewApp()
	w := a.NewWindow("")
	w.SetContent(mv.Container())
	w.ShowAndRun()

	lbl := widget.NewLabel("")
	mv.Push(lbl)

	entry := widget.NewEntry()
	mv.Push(entry)

	mv.Clear()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)
}

func TestMultiView_Replace(t *testing.T) {
	assert := testify.New(t)
	mv := NewMultiView()
	assert.Len(mv.viewStack, 0)
	assert.Len(mv.container.Objects, 0)

	sz := fyne.NewSize(5, 5)
	obj1 := testhelp.NewTestObject(sz)
	obj2 := testhelp.NewTestObject(sz)

	a := test.NewApp()
	w := a.NewWindow("")
	w.SetContent(mv.Container())
	w.ShowAndRun()

	mv.Push(obj1)
	mv.Push(obj2)

	assert.Len(mv.viewStack, 2)
	assert.Len(mv.container.Objects, 1)
	assert.Equal(obj2, mv.container.Objects[0])
	assert.Equal(obj2, mv.viewStack[1])

	obj3 := testhelp.NewTestObject(sz)
	mv.Replace(obj3)

	assert.Len(mv.viewStack, 2)
	assert.Len(mv.container.Objects, 1)
	assert.Equal(obj3, mv.container.Objects[0])
	assert.Equal(obj3, mv.viewStack[1])
}
