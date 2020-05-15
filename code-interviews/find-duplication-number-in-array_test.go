package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicationNumberInArray(t *testing.T) {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	assert.Equal(t, 2, FindDuplicationNumberInArray(arr))

	arr2 := []int{3, 1, 0, 2, 5, 3}
	assert.Equal(t, 3, FindDuplicationNumberInArray(arr2))
}