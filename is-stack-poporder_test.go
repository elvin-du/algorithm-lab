package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsStackPopOrder(t *testing.T) {
	pushSequence := []int{1, 2, 3, 4, 5}
	popSequence := []int{4, 3, 5, 1, 2}
	popSequence2 := []int{4, 5, 3, 2, 1}

	assert.Equal(t, false, IsStackPopOrder(pushSequence, popSequence))
	assert.Equal(t, true, IsStackPopOrder(pushSequence, popSequence2))
}
