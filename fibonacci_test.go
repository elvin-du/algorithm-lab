package algorithm_lab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 13, Fibonacci(6))
}

func TestFibonacci2(t *testing.T) {
	assert.Equal(t, 13, Fibonacci2(6))
}

func TestFibonacci3(t *testing.T) {
	assert.Equal(t, 13, Fibonacci3(6))
}
