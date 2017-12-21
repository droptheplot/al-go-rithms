package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingInteger(t *testing.T) {
	assert.Equal(t, 4, missingInteger([]int{1, 2, 3, 5}))
	assert.Equal(t, 3, missingInteger([]int{1, 2, 4, 5}))
	assert.Equal(t, 2, missingInteger([]int{1, 3, 4, 5}))
}
