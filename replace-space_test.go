package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountSpaceRecursion(t *testing.T) {
	str := ` abc www ddd `
	assert.Equal(t, 4, CountSpaceRecursion(0, str))
	assert.Equal(t, CountSpace(str), CountSpaceRecursion(0, str))
}

func TestReplaceSpace(t *testing.T) {
	str := ` abc www ddd `
	assert.Equal(t, `%20abc%20www%20ddd%20`, ReplaceSpace(str))
}
