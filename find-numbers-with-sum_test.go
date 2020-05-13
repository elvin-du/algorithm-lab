package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindNumbersWithSum(t *testing.T) {
	data := []int{1, 3, 4, 7, 8, 9, 10, 12}
	assert.Equal(t, [2]int([2]int{1, 12}), FindNumbersWithSum(data, 13))
}
