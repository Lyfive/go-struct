/*
	@author: lyfive
	@since: 2023/1/17-17:26
	@desc: //TODO

*
*/
package queue

import "github.com/lyfive/go-struct/list"

type Queue struct {
	list *list.List
}

func New() *Queue {
	return &Queue{list: list.New()}
}

func (q *Queue) Push(data any) {
	q.list.AddToTail(data)
}

func (q *Queue) Pop() any {
	if q.list.Len() == 0 {
		panic("queue is empty!")
		return nil
	}
	return q.list.RemoveHead()
}

func (q *Queue) Empty() bool {
	return q.list.Len() == 0
}

func (q *Queue) Front() any {
	if q.Empty() {
		panic("queue is empty!")
		return nil
	}
	return q.list.Begin().Data
}
