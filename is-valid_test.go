package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValid(t *testing.T) {
	assert.Equal(t, true, isValid("[]"))
	assert.Equal(t, true, isValid("[]{}"))
	assert.Equal(t, true, isValid("{[]}"))
}
