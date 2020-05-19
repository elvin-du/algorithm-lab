package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxInWindows(t *testing.T) {
	data := []int{5, 2, 7, 9, 22, 44, 1}
	assert.Equal(t, []int{7, 9, 22, 44, 44}, MaxInWindows(data, 3))
}