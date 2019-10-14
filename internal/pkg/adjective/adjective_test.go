package adjective

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdjectiveGet tests that the return value is not nil
func TestAdjectiveGet(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(Get(), "should return a value")
}
