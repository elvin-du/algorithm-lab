package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBST(t *testing.T) {
	data := []int{3, 4, 7, 10, 8, 5}
	assert.Equal(t,true,IsBST(data))
}
