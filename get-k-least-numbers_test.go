package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKLeastNumbers(t *testing.T) {
	data := []int{6, 8, 1, 5, 4, 3}
	res := GetKLeastNumbers(data, 3)
	assert.Equal(t, true, AllNumberSmaller(res, 5))
}

func AllNumberSmaller(data []int, target int) bool {
	for _, v := range data {
		if v >= target {
			return false
		}
	}

	return true
}
