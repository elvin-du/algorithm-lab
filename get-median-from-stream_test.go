package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMedian(t *testing.T) {
	InsertStreamData(5)
	InsertStreamData(20)
	assert.Equal(t, 12, GetMedian())

	InsertStreamData(8)
	assert.Equal(t, 8, GetMedian())

	InsertStreamData(2)
	assert.Equal(t, 6, GetMedian())
}
