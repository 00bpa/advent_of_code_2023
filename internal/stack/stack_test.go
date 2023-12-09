package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackLength(t *testing.T) {
	stack := NewStack[int]()
	assert.Equal(t, 0, stack.Length())
	stack.Push(10)
	assert.Equal(t, 1, stack.Length())
	stack.Push(200)
	assert.Equal(t, 2, stack.Length())
	value := stack.Pop()
	assert.Equal(t, 200, value)
	assert.Equal(t, 1, stack.Length())
	value = stack.Pop()
	assert.Equal(t, 10, value)
	assert.Equal(t, 0, stack.Length())
}
