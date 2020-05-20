package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	l, h := 1, 4
	assert.Equal(t, 10, Sum(l, h))
	h = 5
	assert.Equal(t, 15, Sum(l, h))
}

func TestFindContinuousSequence(t *testing.T) {
	res := FindContinuousSequence(100)
	assert.Equal(t, [][]int{[]int{9, 10, 11, 12, 13, 14, 15, 16}, []int{18, 19, 20, 21, 22}}, res)
}
