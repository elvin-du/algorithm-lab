package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintMinNumber(t *testing.T) {
	data := []int{3, 32, 321}
	assert.Equal(t, "321323", PrintMinNumber(data))
}