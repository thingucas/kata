package lc1804_implement_trie_ii_prefix_tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pdt256/kata/go/pkg/lc/leetcode/lc1804_implement_trie_ii_prefix_tree"
)

func TestTrieII(t *testing.T) {
	t.Skip()
	t.Run("example 1", func(t *testing.T) {
		// Given
		trie := lc1804_implement_trie_ii_prefix_tree.New()

		// When
		trie.Insert("apple")
		trie.Insert("apple")

		// Then
		assert.Equal(t, 2, trie.CountWordsEqualTo("apple"))
		assert.Equal(t, 2, trie.CountWordsStartingWith("app"))
		trie.Erase("apple")
		assert.Equal(t, 1, trie.CountWordsEqualTo("apple"))
		assert.Equal(t, 1, trie.CountWordsStartingWith("app"))
		trie.Erase("apple")
		assert.Equal(t, 0, trie.CountWordsStartingWith("app"))
	})
}
