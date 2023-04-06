/*
	@author: lyfive
	@since: 2023/1/17-17:26
	@desc: //TODO
*
*/

package queue

import "github.com/lyfive/go-struct/list"

type Queue[T any] struct {
	list *list.List[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{list: list.New[T]()}
}

func (q *Queue[T]) Push(data T) {
	q.list.AddToTail(data)
}

func (q *Queue[T]) Pop() any {
	if q.list.Len() == 0 {
		panic("queue is empty!")
		return nil
	}
	return q.list.RemoveHead()
}

func (q *Queue[T]) Empty() bool {
	return q.list.Len() == 0
}

func (q *Queue[T]) Front() any {
	if q.Empty() {
		panic("queue is empty!")
		return nil
	}
	return q.list.Begin().Data
}
