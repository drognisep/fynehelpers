package generation

import (
	"errors"
	"testing"

	"fyne.io/fyne/v2/widget"
	testify "github.com/stretchr/testify/require"
)

func TestTreeModelRegistry_AddChild(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()

	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)
	assert.NotEqual("", dataID, "New ID should be returned")
	assert.Equal(reg.Node(dataID), data, "Data should be registered in ID map")
	assert.Contains(reg.Children(ModelRoot), dataID, "Data ID should now appear in child map")
	assert.Equal(ModelRoot, reg.Parent(dataID), "Parent map should contain reference to ModelRoot from dataID")

	data2 := getTreeModelRegistryData()
	_, err = reg.AddChild(dataID, data2)
	assert.NoError(err)
	dataChildren := data.Children()
	assert.Len(dataChildren, 1)
	assert.Equal(data2, dataChildren[0])
}

func TestTreeModelRegistry_AddChild_Neg(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()

	tests := map[string]func(t *testing.T){
		"Non-existent parent ID": func(t *testing.T) {
			assert := testify.New(t)
			_, err := reg.AddChild("1234", data)
			assert.Error(err)
			assert.True(errors.Is(err, ErrNoSuchParent))
		},
		"Nil data": func(t *testing.T) {
			assert := testify.New(t)
			_, err := reg.AddChild(ModelRoot, nil)
			assert.Error(err)
			assert.True(errors.Is(err, ErrNilData))
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}
}

func TestTreeModelRegistry_RemoveChild(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()
	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)

	data2 := getTreeModelRegistryData()
	data2ID, err := reg.AddChild(dataID, data2)
	assert.NoError(err)
	dataChildren := data.Children()
	assert.Len(dataChildren, 1)
	assert.Equal(data2, dataChildren[0])

	reg.RemoveChild(data2ID)
	assert.Len(data.Children(), 0)

	reg.RemoveChild(dataID)
	assert.Nil(reg.Node(dataID), "Data ID should no longer exist in ID map")
	assert.NotContains(reg.Children(ModelRoot), dataID, "Data ID should no longer be listed in ModelRoot's children")
	assert.Nil(reg.Children(ModelRoot), "Ensure that the child map is removed if it's empty")
	_, ok := reg.parentMap[dataID]
	assert.False(ok, "Parent map should no longer contain anything for dataID")
}

func TestTreeModelRegistry_RemoveNonexistentChild(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	dataID := "Something that doesn't exist"
	reg.RemoveChild(dataID)
	assert.Nil(reg.Node(dataID), "Data ID should no longer exist in ID map")
	assert.NotContains(reg.Children(ModelRoot), dataID, "Data ID should no longer be listed in ModelRoot's children")
	assert.Nil(reg.Children(ModelRoot), "Ensure that the child map is removed if it's empty")
	_, ok := reg.parentMap[dataID]
	assert.False(ok, "Parent map should no longer contain anything for dataID")
}

func TestTreeModelRegistry_Children(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()

	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)

	rootChildren := reg.Children(ModelRoot)
	assert.NotNil(rootChildren, "rootChildren should be non-nil")
	assert.Len(rootChildren, 1, "rootChildren should contain dataID")
	assert.Contains(rootChildren, dataID)

	dataChildren := reg.Children(dataID)
	assert.Nil(dataChildren, "Returned child list should be nil")
}

func TestTreeModelRegistry_HasChildren(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()

	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)

	has := reg.HasChildren(ModelRoot)
	assert.True(has, "ModelRoot has a child")

	has = reg.HasChildren(dataID)
	assert.False(has, "Data does not have children")
}

func TestTreeModelRegistry_Node(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()
	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)

	assert.Nil(reg.Node(ModelRoot), "ModelRoot's node should be set to nil")
	assert.Equal(data, reg.Node(dataID), "Data should be in the ID map")
}

func TestTreeModelRegistry_Walk(t *testing.T) {
	assert := testify.New(t)
	reg := NewTreeModelRegistry()
	assert.NotNil(reg)

	data := getTreeModelRegistryData()
	data2 := getTreeModelRegistryData()
	dataID, err := reg.AddChild(ModelRoot, data)
	assert.NoError(err)
	data2ID, err := reg.AddChild(dataID, data2)
	assert.NoError(err)

	rootVisited := 0
	dataVisited := 0
	data2Visited := 0
	reg.Walk(func(parentID widget.TreeNodeID, parent TreeModel, nodeID widget.TreeNodeID, node TreeModel) {
		switch {
		case nodeID == ModelRoot:
			rootVisited++
		case nodeID == dataID:
			dataVisited++
			assert.Equal(data, node)
		case nodeID == data2ID:
			data2Visited++
			assert.Equal(data2, node)
		}
		assert.NotEqual("", nodeID)
	})

	assert.Equal(0, rootVisited, "The root node should never be visited")
	assert.Equal(1, dataVisited, "Data should be visited")
	assert.Equal(1, data2Visited, "Data2 should be visited as well")
}

type ModelData struct {
	BaseTreeModel
	Data string
}

func (d *ModelData) DisplayString() string {
	return d.Data
}

func getTreeModelRegistryData() TreeModel {
	data := ModelData{
		Data: "0",
	}
	return &data
}
