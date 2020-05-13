package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToInt(t *testing.T) {
	str := "-123"
	assert.Equal(t, -123, StrToInt(str))

	str = "0123"
	assert.Equal(t, 123, StrToInt(str))

	str = "0a123"
	assert.Equal(t, 0, StrToInt(str))

	str = "123001"
	assert.Equal(t, 123001, StrToInt(str))
}
