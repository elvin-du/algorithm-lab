package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseSentence(t *testing.T) {
	str := "student. a am I"
	assert.Equal(t, "I am a student.", ReverseSentence(str))
}
