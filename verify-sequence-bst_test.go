package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifySequenceOfBST(t *testing.T) {
	data := []int{5, 7, 6, 9, 12, 10, 8}
	assert.Equal(t, true, VerifySequenceOfBST(data))

	data = []int{5, 7, 99, 9, 12, 10, 8}
	assert.Equal(t, false, VerifySequenceOfBST(data))
}
