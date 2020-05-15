package code_interviews

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstAppearingOnce(t *testing.T) {
	str := "google"
	assert.Equal(t, "l", FirstAppearingOnce(bytes.NewBufferString(str)))
}
