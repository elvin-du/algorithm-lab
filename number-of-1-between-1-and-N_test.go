package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOf1Between1AndN(t *testing.T) {
	assert.Equal(t, 6, NumberOf1Between1AndN(13))
}
