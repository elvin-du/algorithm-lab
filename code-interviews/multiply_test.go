package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiply(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	assert.Equal(t, []int{24, 12, 8, 6}, Multiply(nums))
}
