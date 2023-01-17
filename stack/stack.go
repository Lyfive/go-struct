/*
	@author: lyfive
	@since: 2023/1/17-17:33
	@desc: //TODO

*
*/
package stack

import "github.com/lyfive/go-struct/list"

type Stack struct {
	*list.List
}

func New() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(data any) {
	s.AddToHead(data)
}

func (s *Stack) Pop() any {
	if s.Len() == 0 {
		return nil
	}
	return s.RemoveHead()
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Top() any {
	if s.Empty() {
		return nil
	}
	return s.Begin().Data
}
