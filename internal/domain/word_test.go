package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWordAndGetters(t *testing.T) {
	word := NewWord("test", "category", 1, "hint")

	assert.Equal(t, "test", word.Value())
	assert.Equal(t, "category", word.Category())
	assert.Equal(t, 1, word.Difficulty())
	assert.Equal(t, "hint", word.Hint())
}

func TestWord_Value(t *testing.T) {
	word := NewWord("value", "category", 1, "hint")

	assert.Equal(t, "value", word.Value())
}

func TestWord_Category(t *testing.T) {
	word := NewWord("value", "category", 2, "hint")

	assert.Equal(t, "category", word.Category())
}

func TestWord_Hint(t *testing.T) {
	word := NewWord("value", "category", 2, "hint")

	assert.Equal(t, "hint", word.Hint())
}

func TestWord_Difficulty(t *testing.T) {
	word := NewWord("value", "category", 3, "hint")

	assert.Equal(t, 3, word.Difficulty())
}
