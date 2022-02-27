package lc252_meeting_rooms_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/leetcode/lc252_meeting_rooms"
)

func TestCanAttend(t *testing.T) {
	t.Skip()
	table := []struct {
		Name   string
		Input  [][]int
		Output bool
	}{
		{
			Name: "Example 1",
			Input: [][]int{
				{0, 30},
				{5, 10},
				{15, 20},
			},
			Output: false,
		},
		{
			Name: "Example 2",
			Input: [][]int{
				{7, 10},
				{2, 4},
			},
			Output: true,
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			// Given

			// When
			actual := lc252_meeting_rooms.CanAttend(tt.Input)

			// Then
			assert.Equal(t, tt.Output, actual)
		})
	}
}
