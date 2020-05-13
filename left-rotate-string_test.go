package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inverse(t *testing.T) {
	data := []rune{1, 2, 3}

	inverse(data, 0, len(data)-1)
	assert.Equal(t, []int32{3, 2, 1}, data)
}

func TestLeftRotateString(t *testing.T) {
	str := "123abcd"
	assert.Equal(t, "abcd123", LeftRotateString(str, 3))
}
