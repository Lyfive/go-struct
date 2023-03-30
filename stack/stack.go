/*
	@author: lyfive
	@since: 2023/1/17-17:33
	@desc: //TODO

*
*/
package stack

import "github.com/lyfive/go-struct/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(data any) {
	s.list.AddToHead(data)
}

func (s *Stack) Pop() any {
	if s.list.Len() == 0 {
		panic("stack is empty!")
		return nil
	}
	return s.list.RemoveHead()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

func (s *Stack) Top() any {
	if s.Empty() {
		panic("stack is empty!")
		return nil
	}
	return s.list.Begin().Data
}
