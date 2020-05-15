package code_interviews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexMatch(t *testing.T) {
	str, pattern := `aaa`, `a.a`
	assert.Equal(t, true, RegexMatch(0, 0, str, pattern))
	str2, pattern2 := `aaa`, `ab*ac*a`
	assert.Equal(t, true, RegexMatch(0, 0, str2, pattern2))
	str3, pattern3 := `aaa`, `aa.a`
	assert.Equal(t, false, RegexMatch(0, 0, str3, pattern3))
}