package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, bubbleSort([]int{5, 2, 3, 4, 1}))
	assert.Equal(t, []int{-5, -4, -3, -2, -1}, bubbleSort([]int{-1, -3, -4, -5, -2}))
}
