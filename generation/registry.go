package generation

import (
	"fmt"
	"sync"

	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	ModelRoot widget.TreeNodeID = ""
)

var (
	ErrNoSuchParent = errors.New("no such parent exists")
	ErrNilData      = errors.New("nil data")
)

type modelIdMap = map[widget.TreeNodeID]TreeModel
type modelChildMap = map[widget.TreeNodeID][]widget.TreeNodeID
type modelParentMap = map[widget.TreeNodeID]widget.TreeNodeID

type TreeModelRegistry struct {
	mux       sync.RWMutex
	idMap     modelIdMap
	childMap  modelChildMap
	parentMap modelParentMap
}

func NewTreeModelRegistry() *TreeModelRegistry {
	reg := &TreeModelRegistry{
		idMap:     modelIdMap{},
		childMap:  modelChildMap{},
		parentMap: modelParentMap{},
	}
	reg.idMap[ModelRoot] = nil
	return reg
}

func (r *TreeModelRegistry) AddChild(parentID widget.TreeNodeID, data TreeModel) (widget.TreeNodeID, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	parentNode, ok := r.idMap[parentID]
	if !ok {
		return "", ErrNoSuchParent
	}
	if data == nil {
		return "", ErrNilData
	}
	if err := r.propagateAdd(parentNode, data); err != nil {
		return "", err
	}
	dataID := r.getID()
	r.buildParentLinkage(parentID, data, dataID)
	return dataID, nil
}

func (r *TreeModelRegistry) propagateAdd(parentNode TreeModel, data TreeModel) error {
	if parentNode != nil {
		if err := parentNode.AddChild(data); err != nil {
			return err
		}
	}
	return nil
}

func (r *TreeModelRegistry) buildParentLinkage(parentID widget.TreeNodeID, child TreeModel, childID widget.TreeNodeID) {
	r.idMap[childID] = child
	r.childMap[parentID] = append(r.childMap[parentID], childID)
	r.parentMap[childID] = parentID
	r.buildExtendedLinkage(childID, child)
}

func (r *TreeModelRegistry) buildExtendedLinkage(parentID widget.TreeNodeID, parent TreeModel) {
	for _, c := range parent.Children() {
		cid := r.getID()
		r.buildParentLinkage(parentID, c, cid)
	}
}

func (r *TreeModelRegistry) RemoveChild(dataID widget.TreeNodeID) {
	r.mux.Lock()
	defer r.mux.Unlock()
	parentID, ok := r.parentMap[dataID]
	if !ok {
		return
	}
	r.propagateRemove(r.idMap[parentID], r.idMap[dataID])
	r.tearDownParentLinkage(parentID, dataID)
}

func (r *TreeModelRegistry) propagateRemove(parent TreeModel, child TreeModel) {
	if parent != nil {
		for j, c := range parent.Children() {
			if c == child {
				parent.RemoveChildAt(j)
			}
		}
	}
}

func (r *TreeModelRegistry) tearDownParentLinkage(parentID widget.TreeNodeID, childID widget.TreeNodeID) {
	curChildren := r.childMap[parentID]
	for i, cid := range curChildren {
		if cid == childID {
			r.childMap[parentID] = append(curChildren[:i], curChildren[i+1:]...)
			break
		}
	}
	r.tearDownExtendedLinkage(childID)
	delete(r.idMap, childID)
	if len(r.childMap[parentID]) == 0 {
		delete(r.childMap, parentID)
	}
	delete(r.parentMap, childID)
}

func (r *TreeModelRegistry) tearDownExtendedLinkage(parentID widget.TreeNodeID) {
	for _, cid := range r.childMap[parentID] {
		r.tearDownParentLinkage(parentID, cid)
	}
}

func (r *TreeModelRegistry) Node(nodeID widget.TreeNodeID) TreeModel {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.idMap[nodeID]
}

func (r *TreeModelRegistry) Parent(childID widget.TreeNodeID) widget.TreeNodeID {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.parentMap[childID]
}

func (r *TreeModelRegistry) Children(parentID widget.TreeNodeID) []widget.TreeNodeID {
	r.mux.RLock()
	defer r.mux.RUnlock()
	children := r.childMap[parentID]
	return children
}

func (r *TreeModelRegistry) HasChildren(parentID widget.TreeNodeID) bool {
	r.mux.RLock()
	defer r.mux.RUnlock()
	_, ok := r.childMap[parentID]
	return ok
}

type TreeModelWalkFunc = func(parentID widget.TreeNodeID, parent TreeModel, nodeID widget.TreeNodeID, node TreeModel)

// Walk traverses the registered tree, depth-first. Attempting to modify the tree while walking will result in a deadlock.
func (r *TreeModelRegistry) Walk(walker TreeModelWalkFunc) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	roots := r.Children(ModelRoot)
	for _, child := range roots {
		r.walk(ModelRoot, child, walker)
	}
}

func (r *TreeModelRegistry) walk(parentID, nodeID widget.TreeNodeID, walker TreeModelWalkFunc) {
	node := r.Node(nodeID)
	parent := r.Node(parentID)
	walker(parentID, parent, nodeID, node)
	for _, childID := range r.Children(nodeID) {
		r.walk(nodeID, childID, walker)
	}
}

func (r *TreeModelRegistry) getID() widget.TreeNodeID {
	id, err := uuid.NewRandom()
	for err != nil {
		panic(fmt.Errorf("failed to generate UUID: %v", err))
	}
	return id.String()
}
