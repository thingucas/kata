package meeting_rooms_III_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/google/meeting_rooms_III"
)

func TestCanAttend(t *testing.T) {
	t.Skip()
	table := []struct {
		Name     string
		Meetings [][]int
		Rooms    int
		Queries  [][]int
		Output   []bool
	}{
		{
			Name: "Example 1",
			Meetings: [][]int{
				{1, 2},
				{4, 5},
				{8, 10},
			},
			Rooms: 1,
			Queries: [][]int{
				{2, 3},
				{3, 4},
			},
			Output: []bool{true, true},
		},
		{
			Name: "Example 2",
			Meetings: [][]int{
				{1, 2},
				{4, 5},
				{8, 10},
			},
			Rooms: 1,
			Queries: [][]int{
				{4, 5},
				{5, 6},
			},
			Output: []bool{false, true},
		},
		{
			Name: "Example 3",
			Meetings: [][]int{
				{1, 3},
				{4, 6},
				{6, 8},
				{9, 11},
				{6, 9},
				{1, 3},
				{4, 10},
			},
			Rooms: 3,
			Queries: [][]int{
				{1, 9},
				{2, 6},
				{7, 9},
				{3, 5},
				{3, 9},
				{2, 4},
				{7, 10},
				{5, 9},
				{3, 10},
				{9, 10},
			},
			Output: []bool{false, true, false, true, false, true, false, false, false, true},
		},
	}

	for _, tt := range table {
		t.Run(tt.Name, func(t *testing.T) {
			// Given

			// When
			actual := meeting_rooms_III.CanAdd(tt.Meetings, tt.Rooms, tt.Queries)

			// Then
			assert.Equal(t, tt.Output, actual)
		})
	}
}
