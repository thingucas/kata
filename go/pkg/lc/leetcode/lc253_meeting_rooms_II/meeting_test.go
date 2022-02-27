package lc253_meeting_rooms_II_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/leetcode/lc253_meeting_rooms_II"
)

func TestCanAttend(t *testing.T) {
	t.Skip()
	table := []struct {
		Name   string
		Input  [][]int
		Output int
	}{
		{
			Name: "Example 1",
			Input: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
			},
			Output: 2,
		},
		{
			Name: "Example 2",
			Input: [][]int{
				{7, 10},
				{2, 4},
			},
			Output: 1,
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			// Given

			// When
			actual := lc253_meeting_rooms_II.MinRooms(tt.Input)

			// Then
			assert.Equal(t, tt.Output, actual)
		})
	}
}
