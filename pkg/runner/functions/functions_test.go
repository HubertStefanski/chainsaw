package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFunctions(t *testing.T) {
	assert.Equal(t, 5, len(GetFunctions()))
}
