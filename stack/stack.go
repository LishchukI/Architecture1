package stack

import (
	"fmt"
)

type Element struct {
	data interface{}
	prev *Element
}

type Stack struct {
	size int
	top  *Element
}

func NewStack() Stack {
	return Stack{size: 0}
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(data interface{}) {
	elem := &Element{data, s.top}
	s.top = elem
	s.size++
}

func (s *Stack) Pop() (interface{}, error) {
	elem := s.top
	if elem != nil {
		s.size--
		s.top = elem.prev
		return elem.data, nil
	}
	return nil, fmt.Errorf("Removing item from empty stack")
}
