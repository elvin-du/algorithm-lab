package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstNotRepeatingChar(t *testing.T) {
	str := "abcdeeeeeaaaabbbbwwwwwa"
	assert.Equal(t, 2, FirstNotRepeatingChar(str))
}
