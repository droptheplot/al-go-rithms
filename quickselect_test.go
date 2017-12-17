package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickselect(t *testing.T) {
	assert.Equal(t, 0, quickselect([]int{0}))
	assert.Equal(t, 0, quickselect([]int{2, 0, 1}))
	assert.Equal(t, 0, quickselect([]int{0, 1, 2}))
	assert.Equal(t, 1, quickselect([]int{5, 2, 3, 4, 1}))
	assert.Equal(t, 1, quickselect([]int{6, 2, 3, 1, 4, 6, 6, 5}))
	assert.Equal(t, -5, quickselect([]int{-1, -3, -4, -5, -2}))
}
