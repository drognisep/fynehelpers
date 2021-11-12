package layouthelp

import (
	"testing"

	"fyne.io/fyne/v2"
	testify "github.com/stretchr/testify/require"
)

func TestAccumulateWidth(t *testing.T) {
	tests := map[string]struct {
		sizes []fyne.Size
		expected fyne.Size
	} {
		"Same height": {
			sizes: []fyne.Size{fyne.NewSize(20, 10), fyne.NewSize(20, 10), fyne.NewSize(10, 10)},
			expected: fyne.NewSize(50, 10),
		},
		"Different height": {
			sizes: []fyne.Size{fyne.NewSize(20, 30), fyne.NewSize(20, 20), fyne.NewSize(10, 10)},
			expected: fyne.NewSize(50, 30),
		},
		"No sizes": {
			sizes: []fyne.Size{},
			expected: fyne.NewSize(0, 0),
		},
		"Nil slice": {
			sizes: nil,
			expected: fyne.NewSize(0, 0),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testify.New(t)
			got := AccumulateWidth(tc.sizes...)
			assert.Equal(tc.expected.Width, got.Width)
			assert.Equal(tc.expected.Height, got.Height)
		})
	}
}

func TestAccumulateHeight(t *testing.T) {
	tests := map[string]struct {
		sizes []fyne.Size
		expected fyne.Size
	} {
		"Same width": {
			sizes: []fyne.Size{fyne.NewSize(10, 20), fyne.NewSize(10, 20), fyne.NewSize(10, 10)},
			expected: fyne.NewSize(10, 50),
		},
		"Different width": {
			sizes: []fyne.Size{fyne.NewSize(30, 20), fyne.NewSize(20, 20), fyne.NewSize(10, 10)},
			expected: fyne.NewSize(30, 50),
		},
		"No sizes": {
			sizes: []fyne.Size{},
			expected: fyne.NewSize(0, 0),
		},
		"Nil slice": {
			sizes: nil,
			expected: fyne.NewSize(0, 0),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testify.New(t)
			got := AccumulateHeight(tc.sizes...)
			assert.Equal(tc.expected.Width, got.Width)
			assert.Equal(tc.expected.Height, got.Height)
		})
	}
}
