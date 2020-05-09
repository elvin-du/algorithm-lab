package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	data := []int{5, 2, 7, 9, 22, 44, 1}
	assert.Equal(t, []int{7, 9, 22, 44, 44}, MaxSlidingWindow(data, 3))
}
