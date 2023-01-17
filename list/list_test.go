/*
	@author: lyfive
	@since: 2023/1/17-16:25
	@desc: //TODO

*
*/
package list

import "testing"

func TestList(t *testing.T) {
	ls := New()
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.Push(1)
	ls.Push(2)
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.Reverse()
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.RemoveHead()
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.AddToHead(2)
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.RemoveTail()
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.AddToTail(1)
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
	ls.Push(1)
	ls.Push(2)
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		if i.Data == 1 {
			ls.RemoveNode(i)
		}
	}
	for i := ls.Begin(); i != ls.End(); i = i.Next() {
		t.Log(i.Data)
	}
}
