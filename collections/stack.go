package collections

import (
	"errors"
)

/*
	Stack

	Search O(n)
	Insert O(1)
	Delete O(1)
*/

type StackParams interface {
	int | string | bool
}

type Stack[T StackParams] struct {
	stack []T
	ptr int
	capacity int
}

func NewStack[T StackParams]() *Stack[T] {
	return &Stack[T]{
		stack: []T{},
		ptr: 0,
		capacity: 0,  
	}
}

func (s *Stack[StackParams]) Push(data StackParams){
	*&s.stack = append(*&s.stack, data)
	s.ptr++
	s.capacity++
}	

func (s *Stack[StackParams]) Pop(){
	*&s.stack = (*&s.stack)[:s.ptr-1]
	s.ptr--
}

func (s *Stack[StackParams]) isEmpty() bool{
	if s.ptr == 0 {
		return true
	}

	return false
}

func (s *Stack[StackParams]) Peek() StackParams{
	return *&s.stack[s.ptr-1]
}

func (s *Stack[StackParams]) Clear(){
	*&s.stack = nil
	s.ptr = 0
}

func (s *Stack[StackParams]) GetArr() ([]StackParams,error){
	if s.isEmpty() {
		return []StackParams{}, errors.New("Stack Size is Zero")
	}
	
	var result = make([]StackParams, s.ptr)
	for i, v := range *&s.stack {
		result[i] = v
	}

	return result, nil
}