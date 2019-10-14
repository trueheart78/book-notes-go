package adjective

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdjectiveGet tests that the return value is not nil
func TestAdjectiveGet(t *testing.T) {
	assert.NotNil(t, Get(), "should return a value")
}

// TestAdjectiveWords tests that the return value is not nil
func TestAdjectiveWords(t *testing.T) {
	assert.Equal(t, 47, len(words()), "should match the expect number of words")
}
