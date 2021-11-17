package view

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/drognisep/fynehelpers/generation"
)

var _ fyne.CanvasObject = (*TypeBaseTree)(nil)

// TypeBaseTree is a widget.Tree implementation that manages IDs through generation.TreeModelRegistry.
// This is designed to be the gatekeeper for all widget and model mutations.
type TypeBaseTree struct {
	widget.Tree
	*generation.TreeModelRegistry

	OnTapped          func(id widget.TreeNodeID, model generation.TreeModel, event *fyne.PointEvent) // OnTapped is called by the typeBaseNode that receives an event from Fyne.
	OnDoubleTapped    func(id widget.TreeNodeID, model generation.TreeModel, event *fyne.PointEvent) // OnDoubleTapped is called by the typeBaseNode that receives an event from Fyne.
	OnTappedSecondary func(id widget.TreeNodeID, model generation.TreeModel, event *fyne.PointEvent) // OnTappedSecondary is called by the typeBaseNode that receives an event from Fyne.
}

// NewTypeBaseTree initializes the tree and adds all modelRoots to the registry.
func NewTypeBaseTree(modelRoots ...generation.TreeModel) *TypeBaseTree {
	tree := &TypeBaseTree{
		TreeModelRegistry: generation.NewTreeModelRegistry(),
	}
	tree.Tree = widget.Tree{
		ChildUIDs: tree.Children,
		CreateNode: func(bool) (o fyne.CanvasObject) {
			return newTypeBaseNode(tree)
		},
		IsBranch: tree.HasChildren,
		UpdateNode: func(id widget.TreeNodeID, isBranch bool, node fyne.CanvasObject) {
			treeModel, ok := node.(*typeBaseNode)
			if !ok {
				return
			}
			modelNode := tree.Node(id)
			if modelNode == nil {
				return
			}
			treeModel.update(id, modelNode)
		},
	}
	for _, root := range modelRoots {
		if _, err := tree.AddChild("", root); err != nil {
			log.Printf("Error adding model root: %v\n%v\n", err, root)
		}
	}
	tree.ExtendBaseWidget(tree)
	return tree
}

func (t *TypeBaseTree) Tapped(id widget.TreeNodeID, event *fyne.PointEvent) {
	if t.OnTapped != nil {
		t.OnTapped(id, t.Node(id), event)
	}
}

func (t *TypeBaseTree) DoubleTapped(id widget.TreeNodeID, event *fyne.PointEvent) {
	if t.OnDoubleTapped != nil {
		t.OnDoubleTapped(id, t.Node(id), event)
	}
}

func (t *TypeBaseTree) TappedSecondary(id widget.TreeNodeID, event *fyne.PointEvent) {
	if t.OnTappedSecondary != nil {
		t.OnTappedSecondary(id, t.Node(id), event)
	}
}

func (t *TypeBaseTree) AddChild(parentID widget.TreeNodeID, data generation.TreeModel) (widget.TreeNodeID, error) {
	defer t.Refresh()
	return t.TreeModelRegistry.AddChild(parentID, data)
}

func (t *TypeBaseTree) RemoveChild(dataID widget.TreeNodeID) {
	defer t.Refresh()
	t.TreeModelRegistry.RemoveChild(dataID)
}
