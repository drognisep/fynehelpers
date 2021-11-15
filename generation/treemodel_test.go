package generation

import (
	"strconv"
	"testing"

	"github.com/pkg/errors"
	testify "github.com/stretchr/testify/require"
)

func TestBaseTreeModel_AddChild(t *testing.T) {
	assert := testify.New(t)
	base := &BaseTreeModel{}
	assert.Len(base.children, 0, "No children yet")

	c1 := &BaseTreeModel{}
	c2 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c1))
	assert.NoError(base.AddChild(c2))
	assert.Len(base.children, 2, "Both should be added")
	assert.Equal(c1, base.children[0], "c1 should be the first child")
	assert.Equal(c2, base.children[1], "c2 should be appended")
}

func TestBaseTreeModel_AddChildAt(t *testing.T) {
	base := &BaseTreeModel{}
	tests := []struct {
		Name   string
		Index  int
		Value  *BaseTreeModel
		OldLen int
		NewLen int
	}{
		{
			Name:   "Insert beginning",
			Value:  &BaseTreeModel{},
			Index:  0,
			OldLen: 0,
			NewLen: 1,
		},
		{
			Name:   "Insert end",
			Value:  &BaseTreeModel{},
			Index:  1,
			OldLen: 1,
			NewLen: 2,
		},
		{
			Name:   "Insert middle",
			Value:  &BaseTreeModel{},
			Index:  1,
			OldLen: 2,
			NewLen: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert := testify.New(t)
			assert.Len(base.children, tc.OldLen)
			assert.NoError(base.AddChildAt(tc.Index, tc.Value))
			assert.Len(base.children, tc.NewLen)
			assert.Equal(tc.Value, base.children[tc.Index])
		})
	}
}

func TestBaseTreeModel_AddChildAt_Neg(t *testing.T) {
	base := &BaseTreeModel{}
	c := &BaseTreeModel{}
	if err := base.AddChild(c); err != nil {
		t.Fatal(err)
	}

	tests := map[string]struct {
		Index int
	}{
		"Index < 0": {
			Index: -1,
		},
		"Index > len": {
			Index: 2,
		},
	}

	c2 := &BaseTreeModel{}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testify.New(t)
			err := base.AddChildAt(tc.Index, c2)
			assert.Error(err)
			assert.True(errors.Is(err, ErrBadIndex))
			assert.Len(base.children, 1)
			assert.Equal(base.children[0], c)
		})
	}
}

func TestBaseTreeModel_Children(t *testing.T) {
	assert := testify.New(t)
	base := &BaseTreeModel{}

	c1 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c1))
	c2 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c2))
	c3 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c3))
	assert.Len(base.children, 3)

	children := base.Children()
	assert.Len(children, 3, "Should have the same number of elements")
	assert.True(&children != &base.children, "Should be a list copy")
	assert.Equal(children, base.children, "Contents should still be equal")
}

func TestBaseTreeModel_RemoveChild(t *testing.T) {
	assert := testify.New(t)
	base := &BaseTreeModel{}
	c1 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c1))
	c2 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c2))
	c3 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c3))

	tests := []struct {
		Expected TreeModel
	}{
		{
			Expected: c3,
		},
		{
			Expected: c2,
		},
		{
			Expected: c1,
		},
		{
			Expected: nil,
		},
	}

	for i, tc := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert := testify.New(t)
			removed := base.RemoveChild()
			assert.Equal(tc.Expected, removed)
		})
	}
}

func TestBaseTreeModel_RemoveChildAt(t *testing.T) {
	assert := testify.New(t)
	base := &BaseTreeModel{}
	c1 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c1))
	c2 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c2))
	c3 := &BaseTreeModel{}
	assert.NoError(base.AddChild(c3))

	tests := []struct {
		Name    string
		Index   int
		Removed TreeModel
		OldLen  int
		NewLen  int
	}{
		{
			Name:    "Remove past end",
			Index:   55,
			Removed: nil,
			OldLen:  3,
			NewLen:  3,
		},
		{
			Name:    "Remove past beginning",
			Index:   -7,
			Removed: nil,
			OldLen:  3,
			NewLen:  3,
		},
		{
			Name:    "Remove middle",
			Index:   1,
			Removed: c2,
			OldLen:  3,
			NewLen:  2,
		},
		{
			Name:    "Remove end",
			Index:   1,
			Removed: c3,
			OldLen:  2,
			NewLen:  1,
		},
		{
			Name:    "Remove beginning",
			Index:   0,
			Removed: c1,
			OldLen:  1,
			NewLen:  0,
		},
		{
			Name:    "Remove past end empty",
			Index:   55,
			Removed: nil,
			OldLen:  0,
			NewLen:  0,
		},
		{
			Name:    "Remove past beginning empty",
			Index:   -7,
			Removed: nil,
			OldLen:  0,
			NewLen:  0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			assert := testify.New(t)
			assert.Len(base.children, tc.OldLen)
			removed := base.RemoveChildAt(tc.Index)
			assert.Equal(tc.Removed, removed)
			assert.Len(base.children, tc.NewLen)
		})
	}
}
