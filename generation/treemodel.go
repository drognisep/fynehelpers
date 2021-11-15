package generation

import (
	"sync"

	"fyne.io/fyne/v2"
	"github.com/pkg/errors"
)

// TreeModel must be implemented by any type that should respond to events from a generated tree.
type TreeModel interface {
	DisplayIcon() fyne.Resource      // DisplayIcon returns the icon resource that should be displayed by the tree node. Return nil for no icon.
	DisplayString() string           // DisplayString returns the string that should be displayed by the tree node. Return an empty string for no label.
	Children() []TreeModel           // Children returns a shallow copy of the list of child TreeModel.
	AddChild(TreeModel) error        // AddChild appends a child to the child list. The error value should indicate that the addition was rejected.
	AddChildAt(int, TreeModel) error // AddChildAt adds a child to the child list at a particular location. The error value should indicate that the addition was rejected.
	RemoveChild() TreeModel          // RemoveChild removes a child from the end of the list, if it exists. Returns nil if nothing was removed.
	RemoveChildAt(int) TreeModel     // RemoveChildAt removes a child from the child list if one exists at the given location. Returns nil if nothing was removed or if the index was out of bounds.
}

var _ TreeModel = (*BaseTreeModel)(nil)
var ErrBadIndex = errors.New("invalid index")

// BaseTreeModel is a helper type that implements TreeModel. Users only need to override DisplayIcon and/or DisplayString for read-only or non-persistent models.
type BaseTreeModel struct {
	mux      sync.RWMutex
	children []TreeModel
}

func (b *BaseTreeModel) DisplayIcon() fyne.Resource {
	return nil
}

func (b *BaseTreeModel) DisplayString() string {
	return ""
}

func (b *BaseTreeModel) Children() []TreeModel {
	b.mux.RLock()
	cp := make([]TreeModel, len(b.children))
	copy(cp, b.children)
	b.mux.RUnlock()
	return cp
}

func (b *BaseTreeModel) AddChild(newModel TreeModel) error {
	b.mux.Lock()
	defer b.mux.Unlock()
	b.children = append(b.children, newModel)
	return nil
}

func (b *BaseTreeModel) AddChildAt(index int, newModel TreeModel) error {
	b.mux.Lock()
	defer b.mux.Unlock()

	if index > len(b.children) || index < 0 {
		return errors.Wrapf(ErrBadIndex, "index '%d' out of bounds", index)
	}

	b.children = append(append(b.children[0:index], newModel), b.children[index:]...)
	return nil
}

func (b *BaseTreeModel) RemoveChild() TreeModel {
	b.mux.Lock()
	defer b.mux.Unlock()
	childLen := len(b.children)
	if childLen == 0 {
		return nil
	}
	removed := b.children[childLen-1]
	b.children = b.children[0 : childLen-1]
	return removed
}

func (b *BaseTreeModel) RemoveChildAt(index int) TreeModel {
	b.mux.Lock()
	defer b.mux.Unlock()
	childLen := len(b.children)
	if childLen == 0 {
		return nil
	}
	if index < 0 || index > childLen {
		return nil
	}
	removed := b.children[index]
	b.children = append(b.children[0:index], b.children[index+1:]...)
	return removed
}
