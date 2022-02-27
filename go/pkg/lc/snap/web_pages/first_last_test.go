package web_pages_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pdt256/kata/go/pkg/lc/snap/web_pages"
)

func TestFirstLast(t *testing.T) {
	t.Skip()
	// Given
	table := []struct {
		Name     string
		Input    [][]string
		Expected [][]string
	}{
		{
			Name: "in order",
			Input: [][]string{
				{"page1", "page2"},
				{"page2", "page3"},
				{"page3", "page4"},
				{"page4", "page5"},
			},
			Expected: [][]string{
				{"page1", "page5"},
			},
		},
		{
			Name: "shuffled",
			Input: [][]string{
				{"page2", "page3"},
				{"page1", "page2"},
				{"page4", "page5"},
				{"page3", "page4"},
			},
			Expected: [][]string{
				{"page1", "page5"},
			},
		},
		{
			Name: "with cycle",
			Input: [][]string{
				{"page2", "page3"},
				{"page1", "page2"},
				{"page4", "page5"},
				{"page3", "page4"},
				{"page2", "page1"},
			},
			Expected: [][]string{
				{"page1", "page5"},
				{"page2", "page5"},
			},
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			// Given

			// When
			output := web_pages.FirstLast(tt.Input)

			// Then
			total := 0
			require.NotNil(t, output)
			for _, expected := range tt.Expected {
				if expected[0] == output[0] && expected[1] == output[1] {
					total++
				}
			}
			assert.Equal(t, 1, total)
		})
	}
}
