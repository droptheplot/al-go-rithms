package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 2, binarySearch([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, 0, binarySearch([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal(t, 4, binarySearch([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal(t, -1, binarySearch([]int{1, 2, 3, 4, 5}, 0))
}
