package shortest_land_path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/snap/shortest_land_path"
)

func TestShortestPath(t *testing.T) {
	t.Skip()
	t.Run("", func(t *testing.T) {
		// Given
		input := [][]int{
			{1, 0, 0, 1, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 1, 0, 0},
			{0, 1, 1, 0, 1},
			{0, 1, 0, 0, 0},
		}

		expectedOutput := [][]int{
			{0, 0, 0, 1, 0},
			{0, 0, 1, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 1, 1, 0, 0},
			{0, 1, 0, 0, 0},
		}
		// When
		output := shortest_land_path.ShortestPath(input)

		// Then
		assert.Equal(t, expectedOutput, output)
	})
}
