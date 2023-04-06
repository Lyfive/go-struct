/*
	@author: lyfive
	@since: 2023/3/30-21:19
	@desc: //DONE

*
*/
package stack

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	s.Push(123)
	//s.Push("strings")
	s.Push(456)
	s1 := s.Top()
	s2 := s.Pop()
	s3 := s.Pop()
	fmt.Println(s1, s2, s3)
	fmt.Println(s.Empty())
}

func TestNewAny(t *testing.T) {
	s := New[any]()
	s.Push(123)
	s.Push("strings")
	s.Push(456)
	s1 := s.Top()
	s2 := s.Pop()
	s3 := s.Pop()
	fmt.Println(s1, s2, s3)
	fmt.Println(s.Empty())
}
