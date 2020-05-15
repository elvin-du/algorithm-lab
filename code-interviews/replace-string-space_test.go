package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceStringSpace(t *testing.T) {
	str := ` abc www ddd `
	assert.Equal(t, `%20abc%20www%20ddd%20`, ReplaceStringSpace(str))
}