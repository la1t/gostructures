package stack

import (
	"container/list"
	"errors"
)

type Stack struct {
	dl *list.List
}

func New() *Stack {
	dl := list.New()
	return &Stack{dl: dl}
}

func (s *Stack) Push(value int) {
	s.dl.PushBack(value)
}

func (s *Stack) Pop() (int, error) {
	last := s.dl.Back()
	if last == nil {
		return 0, errors.New("stack is empty")
	}
	s.dl.Remove(last)
	return last.Value.(int), nil
}
