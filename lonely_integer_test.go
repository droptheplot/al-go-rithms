package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLonelyInteger(t *testing.T) {
	assert.Equal(t, 3, lonelyInteger([]int{1, 1, 2, 2, 3, 4, 4, 5, 5}))
	assert.Equal(t, 1, lonelyInteger([]int{1, 2, 2, 3, 3, 4, 4, 5, 5}))
	assert.Equal(t, 5, lonelyInteger([]int{1, 1, 2, 2, 3, 3, 4, 4, 5}))
}
