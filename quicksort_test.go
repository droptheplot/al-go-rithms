package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, quicksort([]int{2, 0, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, quicksort([]int{5, 2, 3, 4, 1}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 6, 6}, quicksort([]int{6, 2, 3, 1, 4, 6, 6, 5}))
	assert.Equal(t, []int{-5, -4, -3, -2, -1}, quicksort([]int{-1, -3, -4, -5, -2}))
}
