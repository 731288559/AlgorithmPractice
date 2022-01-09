package algorithm

import "container/list"

type Stack struct {
	l *list.List
}

func NewStack() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Pop() int {
	e := s.l.Back()
	if e != nil {
		s.l.Remove(e)
		return e.Value.(int)
	}
	return -1
}

func (s *Stack) Push(ele int) {
	s.l.PushBack(ele)
}

func (s *Stack) Size() int {
	return s.l.Len()
}

func (s *Stack) Top() int {
	e := s.l.Back()
	if e != nil {
		return e.Value.(int)
	}
	return -1
}

func StackTest() {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	println(s.Pop())
	println(s.Pop())
}
