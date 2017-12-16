package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {
	assert.Equal(t, 2, linearSearch([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, 0, linearSearch([]int{1, 2, 3, 4, 5}, 1))
	assert.Equal(t, 4, linearSearch([]int{1, 2, 3, 4, 5}, 5))
	assert.Equal(t, -1, linearSearch([]int{1, 2, 3, 4, 5}, 0))
}
