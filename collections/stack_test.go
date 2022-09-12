package collections

import (
	"log"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_stack_init(t *testing.T) {
	s := NewStack[int]()
	
	len := s.ptr
	capacity := s.capacity
	
	assert.Equal(t, len,0 )
	assert.Equal(t, capacity,0 )
}

func Test_Push(t *testing.T){
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	result, err:= s.GetArr()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, result, []int{1,2,3,4})
	assert.Equal(t, s.ptr, 4)
}

func Test_Pop(t *testing.T){
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Pop()
	s.Pop()

	result, err:= s.GetArr()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.Equal(t, result, []int{1,2})
	assert.Equal(t, s.ptr, 2)
}

func Test_peek(t *testing.T){
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	result := s.Peek()
	assert.Equal(t, result, 4)
	assert.Equal(t, s.ptr, 4)
}

func Test_Clear(t *testing.T){
	s := NewStack[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	s.Clear()
	assert.Equal(t, s.ptr , 0)
}