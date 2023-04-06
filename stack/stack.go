/*
	@author: lyfive
	@since: 2023/1/17-17:33
	@desc: //TODO

*
*/
package stack

import "github.com/lyfive/go-struct/list"

type Stack[T any] struct {
	list *list.List[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{list.New[T]()}
}

func (s *Stack[T]) Push(data T) {
	s.list.AddToHead(data)
}

func (s *Stack[T]) Pop() T {
	if s.list.Len() == 0 {
		panic("stack is empty!")
	}
	return s.list.RemoveHead()
}

func (s *Stack[T]) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack[T]) Top() T {
	if s.Empty() {
		panic("stack is empty!")
	}
	return s.list.Begin().Data
}
