package lc505_maze_II_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/leetcode/lc505_maze_II"
)

func TestShortestDistance(t *testing.T) {
	t.Skip()
	t.Run("example 1", func(t *testing.T) {
		// Given
		maze := [][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}
		start := []int{0, 4}
		end := []int{4, 4}

		// When
		actual := lc505_maze_II.ShortestDistance(maze, start, end)

		// Then
		assert.Equal(t, 12, actual)
	})

	t.Run("example 2", func(t *testing.T) {
		// Given
		maze := [][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		}
		start := []int{0, 4}
		end := []int{3, 2}

		// When
		actual := lc505_maze_II.ShortestDistance(maze, start, end)

		// Then
		assert.Equal(t, -1, actual)
	})
}
