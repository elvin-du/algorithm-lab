package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoreThanHalfNumber(t *testing.T) {
	data := []int{2, 2, 2, 3, 2, 5, 7}
	assert.Equal(t, 2, MoreThanHalfNumber(data))

	data = []int{2, 2, 2, 3, 2, 5, 7, 8}
	assert.Equal(t, -1, MoreThanHalfNumber(data))

	data = []int{1, 2, 2, 2, 3, 2, 5, 7, 8}
	assert.Equal(t, -1, MoreThanHalfNumber2(data))
}
