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

type Node struct {
	Data any
	next *Node
	pre  *Node
	list *List
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) Pre() *Node {
	return node.pre
}

// List is a bidirectional list, head and tail just for convenience
type List struct {
	head *Node
	tail *Node
	cnt  int
}

// New return a new list
func New() *List {
	newList := &List{
		head: &Node{},
		tail: &Node{},
		cnt:  0,
	}
	newList.head.next = newList.tail
	newList.head.pre = nil
	newList.tail.pre = newList.head
	newList.tail.next = nil
	return newList
}

// AddToHead add a node to head
func (list *List) AddToHead(data any) {
	newNode := &Node{
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
func (list *List) AddToTail(data any) {
	newNode := &Node{
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
func (list *List) Push(data any) {
	list.AddToTail(data)
}

// AddNode add a node after a node
func (list *List) AddNode(node *Node, data any) error {
	if node == nil {
		return errors.New("node is nil")
	}
	newNode := &Node{
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
func (list *List) RemoveHead() any {
	if list.head.next == list.tail {
		return nil
	}
	data := list.head.next.Data
	list.head.list = nil
	list.head.next = list.head.next.next
	list.head.next.pre = list.head
	list.cnt--
	return data
}

// RemoveTail remove tail node
func (list *List) RemoveTail() any {
	if list.tail.pre == list.head {
		return nil
	}
	data := list.tail.pre.Data
	list.tail.pre.list = nil
	list.tail.pre = list.tail.pre.pre
	list.tail.pre.next = list.tail
	list.cnt--
	return data
}

// Pop tail node
func (list *List) Pop() any {
	return list.RemoveTail()
}

// RemoveNode remove a node
func (list *List) RemoveNode(node *Node) error {
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
func (list *List) RemoveData(data any) {
	for node := list.head.next; node != list.tail; node = node.next {
		if node.Data == data {
			_ = list.RemoveNode(node)
			return
		}
	}
}

// FindNode find a node is in the list
func (list *List) FindNode(node *Node) bool {
	return node.list == list
}

// FindData find a node by data
func (list *List) FindData(data any) *Node {
	for node := list.head.next; node != list.tail; node = node.next {
		if node.Data == data {
			return node
		}
	}
	return nil
}

// Reverse the list
func (list *List) Reverse() {
	for node := list.head; node != nil; node = node.pre {
		node.next, node.pre = node.pre, node.next
	}
	list.head, list.tail = list.tail, list.head
}

// Begin return the list begin pointer
func (list *List) Begin() *Node {
	return list.head.next
}

// End return the list end pointer
func (list *List) End() *Node {
	return list.tail
}

// Len return this list length
func (list *List) Len() int {
	return list.cnt
}
