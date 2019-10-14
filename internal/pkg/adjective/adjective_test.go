package adjective

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSample description
func TestSample(t *testing.T) {
	assert := assert.New(t)
	// assert equality
	assert.NotNil(Get(), "should return a value")
}
