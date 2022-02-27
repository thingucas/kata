package visit_friends_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/snap/visit_friends"
)

func TestFindShortestPath(t *testing.T) {
	t.Skip()
	t.Run("", func(t *testing.T) {
		// Given
		input := ToGrid([]string{
			"..P..",
			"F...F",
			"FF.FF",
		})

		// When
		res := visit_friends.FindShortestPath(input)

		// Then
		assert.Equal(t, 9, res)
	})
}

func ToGrid(input []string) [][]byte {
	grid := make([][]byte, len(input))
	for row, line := range input {
		grid[row] = make([]byte, len(line))
		for col, c := range line {
			grid[row][col] = byte(c)
		}
	}
	return grid
}
