package primefactors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/primefactors"
)

func TestPrimeFactors(t *testing.T) {
	assert.Equal(t, []int(nil), primefactors.Generate(1))
	// assert.Equal(t, []int{2}, primefactors.Generate(2))
}

func TestAllPrimeFactors(t *testing.T) {
	t.Skip()
	for _, data := range primeFactorsData {
		assert.Equal(t, data.expected, primefactors.Generate(data.number))
	}
}

var primeFactorsData = []struct {
	number   int
	expected []int
}{
	{1, []int(nil)},
	{2, []int{2}},
	{3, []int{3}},
	{4, []int{2, 2}},
	{5, []int{5}},
	{6, []int{2, 3}},
	{7, []int{7}},
	{8, []int{2, 2, 2}},
	{9, []int{3, 3}},
	{360, []int{2, 2, 2, 3, 3, 5}},
	{2 * 3 * 5 * 7 * 13, []int{2, 3, 5, 7, 13}},
}
