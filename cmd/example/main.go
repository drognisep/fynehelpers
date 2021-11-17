package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/drognisep/fynehelpers/cmd/example/view"
	"github.com/drognisep/fynehelpers/generation"
)

func main() {
	w := app.New().NewWindow("Tree test")
	w.Resize(fyne.NewSize(640, 480))
	w.CenterOnScreen()

	data := &TestData{
		Data: "root",
	}
	_ = data.AddChild(&TestData{
		Data: "A",
	})
	data2 := &TestData{
		Data: "root2",
	}
	_ = data2.AddChild(&TestData{
		Data: "B",
	})
	tree := view.NewTypeBaseTree(data, data2)
	var i int
	tree.OnDoubleTapped = func(id widget.TreeNodeID, model generation.TreeModel, event *fyne.PointEvent) {
		log.Printf("Adding child to parent '%s'\n", id)
		_, _ = tree.AddChild(id, &TestData{
			Data: fmt.Sprintf("%s_%d", model.DisplayString(), i),
		})
		i++
		log.Println("Done!")
	}
	tree.OnTappedSecondary = func(id widget.TreeNodeID, model generation.TreeModel, event *fyne.PointEvent) {
		tree.RemoveChild(id)
	}
	w.SetContent(tree)
	w.ShowAndRun()
}

var _ generation.TreeModel = (*TestData)(nil)

type TestData struct {
	generation.BaseTreeModel

	Data string
}

func (t *TestData) DisplayString() string {
	return t.Data
}

func (t *TestData) DisplayIcon() fyne.Resource {
	return theme.CheckButtonCheckedIcon()
}
