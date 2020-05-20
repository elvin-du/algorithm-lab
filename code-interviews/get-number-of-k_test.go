package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNumberOfK(t *testing.T) {
	data := []int{1, 3, 5, 7, 7, 8, 8, 8, 10, 22, 30}
	assert.Equal(t, 3, GetNumberOfK(data, 8))
	assert.Equal(t, 2, GetNumberOfK(data, 7))
	assert.Equal(t, 1, GetNumberOfK(data, 3))
}