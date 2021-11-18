package layouthelp

import (
	"testing"

	"fyne.io/fyne/v2"
	"github.com/drognisep/fynehelpers/testhelp"
	testify "github.com/stretchr/testify/require"
)

func TestPaddingLayout_MinSize(t *testing.T) {
	sz := fyne.NewSize(5, 5)
	obj1 := testhelp.NewTestObject(sz)
	obj2 := testhelp.NewTestObject(sz)

	tests := map[string]struct {
		lyt      fyne.Layout
		objs     []fyne.CanvasObject
		expected fyne.Size
	}{
		"Single object": {
			lyt:      new(marginLayout),
			objs:     []fyne.CanvasObject{obj1},
			expected: obj1.MinSize(),
		},
		"Multiple objects": {
			lyt:      new(marginLayout),
			objs:     []fyne.CanvasObject{obj1, obj2},
			expected: obj1.MinSize(),
		},
		"Single object with margin": {
			lyt: &marginLayout{
				top:    5,
				bottom: 5,
				left:   5,
				right:  5,
			},
			objs:     []fyne.CanvasObject{obj1},
			expected: obj1.MinSize().Add(fyne.NewSize(10, 10)),
		},
		"Multiple objects with margin": {
			lyt: &marginLayout{
				top:    5,
				bottom: 5,
				left:   5,
				right:  5,
			},
			objs:     []fyne.CanvasObject{obj1, obj2},
			expected: obj1.MinSize().Add(fyne.NewSize(10, 10)),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testify.New(t)
			sz := tc.lyt.MinSize(tc.objs)
			assert.Equal(tc.expected.Width, sz.Width)
			assert.Equal(tc.expected.Height, sz.Height)
		})
	}
}

func TestPaddingLayout_Layout(t *testing.T) {
	sz := fyne.NewSize(5, 5)
	obj1 := testhelp.NewTestObject(sz)
	obj2 := testhelp.NewTestObject(sz)

	tests := map[string]struct {
		lyt         fyne.Layout
		objs        []fyne.CanvasObject
		expectedSz  fyne.Size
		expectedPos fyne.Position
	}{
		"Single object": {
			lyt:         new(marginLayout),
			objs:        []fyne.CanvasObject{obj1},
			expectedSz:  obj1.MinSize(),
			expectedPos: fyne.Position{},
		},
		"Multiple objects": {
			lyt:         new(marginLayout),
			objs:        []fyne.CanvasObject{obj1, obj2},
			expectedSz:  obj1.MinSize(),
			expectedPos: fyne.Position{},
		},
		"Single object with margin": {
			lyt: &marginLayout{
				top:    5,
				bottom: 5,
				left:   5,
				right:  5,
			},
			objs:        []fyne.CanvasObject{obj1},
			expectedSz:  obj1.MinSize(),
			expectedPos: fyne.NewPos(5, 5),
		},
		"Multiple objects with margin": {
			lyt: &marginLayout{
				top:    5,
				bottom: 5,
				left:   5,
				right:  5,
			},
			objs:        []fyne.CanvasObject{obj1, obj2},
			expectedSz:  obj1.MinSize(),
			expectedPos: fyne.NewPos(5, 5),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testify.New(t)
			assert.NotPanics(func() {
				tc.lyt.Layout(tc.objs, fyne.NewSize(100, 100))
			})
			objSz := tc.objs[0].Size()
			objPs := tc.objs[0].Position()
			assert.Equal(tc.expectedSz.Width, objSz.Width)
			assert.Equal(tc.expectedSz.Height, objSz.Height)
			assert.Equal(tc.expectedPos.X, objPs.X)
			assert.Equal(tc.expectedPos.Y, objPs.Y)
		})
	}
}
