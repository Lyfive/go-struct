/*
	@author: lyfive
	@since: 2023/1/17-17:26
	@desc: //TODO

*
*/
package queue

import "github.com/lyfive/go-struct/list"

type Queue struct {
	*list.List
}

func New() *Queue {
	return &Queue{list.New()}
}

func (q *Queue) Push(data any) {
	q.AddToTail(data)
}

func (q *Queue) Pop() any {
	if q.Len() == 0 {
		return nil
	}
	return q.RemoveHead()
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Front() any {
	if q.Empty() {
		return nil
	}
	return q.Begin().Data
}
