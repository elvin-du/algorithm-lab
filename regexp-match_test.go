package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegexpMatch(t *testing.T) {
	str, pattern := `aaa`, `a.a`
	assert.Equal(t, true, RegexpMatch(0, 0, str, pattern))
	str2, pattern2 := `aaa`, `ab*ac*a`
	assert.Equal(t, true, RegexpMatch(0, 0, str2, pattern2))
	str3, pattern3 := `aaa`, `aa.a`
	assert.Equal(t, false, RegexpMatch(0, 0, str3, pattern3))
}
