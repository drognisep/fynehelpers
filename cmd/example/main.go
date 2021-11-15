package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"github.com/drognisep/fynehelpers/cmd/example/view"
	"github.com/drognisep/fynehelpers/generation"
)

func main() {
	w := app.New().NewWindow("Tree test")
	w.CenterOnScreen()

	data := &TestData{}
	node := view.NewTypeBaseNode()
	w.SetContent(node)
	time.AfterFunc(time.Second, func() {
		view.Update(node, data)
	})
	w.ShowAndRun()
}

var _ generation.TreeModel = (*TestData)(nil)

type TestData struct {
	generation.BaseTreeModel
}

func (t *TestData) DisplayString() string {
	return "Testing"
}

func (t *TestData) DisplayIcon() fyne.Resource {
	return theme.CheckButtonCheckedIcon()
}
