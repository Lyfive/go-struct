/*
	@author: lyfive
	@since: 2023/3/30-21:15
	@desc: //DONE

*
*/
package queue

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	q := New()
	q.Push(123)
	q.Push("strings")
	fmt.Println(q.Empty())
	q1 := q.Front()
	q2 := q.Pop()
	q3 := q.Pop()
	fmt.Println(q1, q2, q3)
	fmt.Println(q.Empty())
}
