package Router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestFunction(t *testing.T) {
	result := test(3, 4)
	assert.Equal(t, result, 7)
}
