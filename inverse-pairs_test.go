package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInversePairs(t *testing.T) {
	data := []int{6, 8, 2, 4}
	InversePairs(data)
	assert.Equal(t, []int{2, 4, 6, 8}, data)
	assert.Equal(t, 4, pairsCnt)
}
