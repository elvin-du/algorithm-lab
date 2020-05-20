package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_inverse(t *testing.T) {
	data := []rune{1, 2, 3}

	rotate(data)
	assert.Equal(t, []rune{3, 2, 1}, data)
}

func TestLeftRotateString(t *testing.T) {
	str := "123abcd"
	assert.Equal(t, "abcd123", LeftRotateString(str, 3))
}
