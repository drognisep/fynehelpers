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

func (r *TreeModelRegistry) AddChild(parent widget.TreeNodeID, data TreeModel) (widget.TreeNodeID, error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	parentNode, ok := r.idMap[parent]
	if !ok {
		return "", ErrNoSuchParent
	}
	if data == nil {
		return "", ErrNilData
	}
	if parentNode != nil {
		if err := parentNode.AddChild(data); err != nil {
			return "", err
		}
	}
	id := r.getID()
	r.idMap[id] = data
	r.childMap[parent] = append(r.childMap[parent], id)
	r.parentMap[id] = parent
	return id, nil
}

func (r *TreeModelRegistry) RemoveChild(dataID widget.TreeNodeID) {
	r.mux.Lock()
	defer r.mux.Unlock()
	parentID, ok := r.parentMap[dataID]
	if !ok {
		return
	}
	parent := r.idMap[parentID]
	curChildren := r.childMap[parentID]
	for i, cid := range curChildren {
		if cid == dataID {
			if parent != nil {
				child := r.idMap[cid]
				for j, c := range parent.Children() {
					if c == child {
						parent.RemoveChildAt(j)
					}
				}
			}
			r.childMap[parentID] = append(curChildren[:i], curChildren[i+1:]...)
			break
		}
	}
	delete(r.idMap, dataID)
	if len(r.childMap[parentID]) == 0 {
		delete(r.childMap, parentID)
	}
	delete(r.parentMap, dataID)
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
