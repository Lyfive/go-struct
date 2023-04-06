/*
	@author: lyfive
	@since: 2023/1/17-16:11
	@desc: //TODO
*
*/

package list

import (
	"errors"
)

type Node[T any] struct {
	Data T
	next *Node[T]
	pre  *Node[T]
	list *List[T]
}

func (node *Node[T]) Next() *Node[T] {
	return node.next
}

func (node *Node[T]) Pre() *Node[T] {
	return node.pre
}

// List is a bidirectional list, head and tail just for convenience
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	cnt  int
}

// New return a new list
func New[T any]() *List[T] {
	newList := &List[T]{
		head: &Node[T]{},
		tail: &Node[T]{},
		cnt:  0,
	}
	newList.head.next = newList.tail
	newList.head.pre = nil
	newList.tail.pre = newList.head
	newList.tail.next = nil
	return newList
}

// AddToHead add a node to head
func (list *List[T]) AddToHead(data T) {
	newNode := &Node[T]{
		Data: data,
		list: list,
	}
	// from head <-> headNext
	// to head <-> newNode <-> headNext
	newNode.next = list.head.next
	newNode.pre = list.head
	list.head.next.pre = newNode
	list.head.next = newNode
	list.cnt++
}

// AddToTail add a node to tail
func (list *List[T]) AddToTail(data T) {
	newNode := &Node[T]{
		Data: data,
		list: list,
	}
	// from tailPre <-> tail
	// to tailPre <-> newNode <-> tail
	newNode.next = list.tail
	newNode.pre = list.tail.pre
	list.tail.pre.next = newNode
	list.tail.pre = newNode
	list.cnt++
}

// Push a node to list
func (list *List[T]) Push(data T) {
	list.AddToTail(data)
}

// AddNode add a node after a node
func (list *List[T]) AddNode(node *Node[T], data T) error {
	if node == nil {
		return errors.New("node is nil")
	}
	newNode := &Node[T]{
		Data: data,
		list: list,
	}
	newNode.next = node.next
	newNode.pre = node
	node.next.pre = newNode
	node.next = newNode
	list.cnt++
	return nil
}

// RemoveHead remove head node
func (list *List[T]) RemoveHead() T {
	if list.head.next == list.tail {
		panic("List is empty!")
		//return nil
	}
	data := list.head.next.Data
	list.head.list = nil
	list.head.next = list.head.next.next
	list.head.next.pre = list.head
	list.cnt--
	return data
}

// RemoveTail remove tail node
func (list *List[T]) RemoveTail() T {
	if list.tail.pre == list.head {
		panic("List is empty!")
		//return nil
	}
	data := list.tail.pre.Data
	list.tail.pre.list = nil
	list.tail.pre = list.tail.pre.pre
	list.tail.pre.next = list.tail
	list.cnt--
	return data
}

// Pop tail node
func (list *List[T]) Pop() any {
	return list.RemoveTail()
}

// RemoveNode remove a node
// Remove but next and pre can be used
func (list *List[T]) RemoveNode(node *Node[T]) error {
	if node == nil {
		return errors.New("node is nil")
	} else if node.list != list {
		return errors.New("node is not in this list")
	}
	node.list = nil
	node.pre.next = node.next
	node.next.pre = node.pre
	list.cnt--
	return nil
}

// RemoveData remove a node by data
//func (list *List[T]) RemoveData(data T) {
//    for node := list.head.next; node != list.tail; node = node.next {
//        if node.Data == data {
//            _ = list.RemoveNode(node)
//            return
//        }
//    }
//}

// FindNode find a node is in the list
func (list *List[T]) FindNode(node *Node[T]) bool {
	return node.list == list
}

// FindData find a node by data
//func (list *List) FindData(data T) *Node {
//    for node := list.head.next; node != list.tail; node = node.next {
//        if node.Data == data {
//            return node
//        }
//    }
//    return nil
//}

// Reverse the list
func (list *List[T]) Reverse() {
	for node := list.head; node != nil; node = node.pre {
		node.next, node.pre = node.pre, node.next
	}
	list.head, list.tail = list.tail, list.head
}

// Begin return the list begin pointer
func (list *List[T]) Begin() *Node[T] {
	return list.head.next
}

// End return the list end pointer
func (list *List[T]) End() *Node[T] {
	return list.tail
}

// Len return this list length
func (list *List[T]) Len() int {
	return list.cnt
}
