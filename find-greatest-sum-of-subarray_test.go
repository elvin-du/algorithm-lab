package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindGreatestSumOfSubArray(t *testing.T) {
	data := []int{6, -3, -2, 7, -15, 1, 2, 2}
	assert.Equal(t, 8, FindGreatestSumOfSubArray(data))
}
