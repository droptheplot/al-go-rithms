package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, uint64(1), fibonacci(1))
	assert.Equal(t, uint64(1), fibonacci(2))
	assert.Equal(t, uint64(5), fibonacci(5))
	assert.Equal(t, uint64(55), fibonacci(10))
	assert.Equal(t, uint64(12586269025), fibonacci(50))
}
