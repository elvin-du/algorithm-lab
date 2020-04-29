package algorithm_lab

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstAppearingOnce(t *testing.T) {
	r := bytes.NewBufferString("google")
	assert.EqualValues(t, 'l', FirstAppearingOnce(r))

	r2 := bytes.NewBufferString("go")
	assert.EqualValues(t, 'g', FirstAppearingOnce(r2))
}
