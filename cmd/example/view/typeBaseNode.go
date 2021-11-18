package view

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/drognisep/fynehelpers/generation"
	"github.com/drognisep/fynehelpers/layouthelp"
)

var _ fyne.Widget = (*typeBaseNode)(nil)
var _ fyne.Tappable = (*typeBaseNode)(nil)
var _ fyne.DoubleTappable = (*typeBaseNode)(nil)
var _ fyne.SecondaryTappable = (*typeBaseNode)(nil)

type typeBaseNode struct {
	widget.BaseWidget

	mux    sync.RWMutex
	id     widget.TreeNodeID
	render *typeBaseNodeRenderer
	tree   *TypeBaseTree
}

func (t *typeBaseNode) CreateRenderer() fyne.WidgetRenderer {
	t.render = newTypeBaseNodeRenderer()
	return t.render
}

func newTypeBaseNode(tree *TypeBaseTree) *typeBaseNode {
	node := &typeBaseNode{
		tree: tree,
	}
	node.ExtendBaseWidget(node)
	return node
}

func (t *typeBaseNode) update(id widget.TreeNodeID, model generation.TreeModel) {
	t.mux.Lock()
	t.id = id
	if t.render != nil {
		t.render.label.SetText(model.DisplayString())
		t.render.icon.SetResource(model.DisplayIcon())
	}
	t.mux.Unlock()
	t.Refresh()
}

func (t *typeBaseNode) Tapped(event *fyne.PointEvent) {
	t.mux.RLock()
	id := t.id
	t.mux.RUnlock()
	t.tree.Tapped(id, event)
}

func (t *typeBaseNode) DoubleTapped(event *fyne.PointEvent) {
	t.mux.RLock()
	id := t.id
	t.mux.RUnlock()
	t.tree.DoubleTapped(id, event)
}

func (t *typeBaseNode) TappedSecondary(event *fyne.PointEvent) {
	t.mux.RLock()
	id := t.id
	t.mux.RUnlock()
	t.tree.TappedSecondary(id, event)
}

var _ fyne.WidgetRenderer = (*typeBaseNodeRenderer)(nil)

type typeBaseNodeRenderer struct {
	icon    *widget.Icon
	label   *widget.Label
	layout  fyne.Layout
	objects []fyne.CanvasObject
}

func newTypeBaseNodeRenderer() *typeBaseNodeRenderer {
	render := &typeBaseNodeRenderer{
		icon: &widget.Icon{},
		label: &widget.Label{
			Alignment: fyne.TextAlignLeading,
			TextStyle: fyne.TextStyle{},
		},
		layout: layout.NewHBoxLayout(),
	}
	render.objects = []fyne.CanvasObject{render.icon, render.label}
	return render
}

func (r *typeBaseNodeRenderer) Destroy() {
	r.icon = nil
	r.label = nil
	r.layout = nil
	r.objects = nil
}

func (r *typeBaseNodeRenderer) Layout(parent fyne.Size) {
	r.layout.Layout(r.objects, parent)
}

func (r *typeBaseNodeRenderer) MinSize() fyne.Size {
	return layouthelp.AccumulateWidth(r.icon.MinSize(), r.label.MinSize())
}

func (r *typeBaseNodeRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *typeBaseNodeRenderer) Refresh() {
	for _, obj := range r.objects {
		obj.Refresh()
	}
}
