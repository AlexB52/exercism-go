package linkedlist

import (
	"errors"
)

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type List struct {
	first, last *Node
}

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

var ErrEmptyList = errors.New("list is empty")

func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{}
	}

	var nodes = make([]*Node, len(args))
	nodes[0] = &Node{Value: args[0]}
	for i := 1; i < len(args); i++ {
		nodes[i] = &Node{Value: args[i]}
		nodes[i].prev, nodes[i-1].next = nodes[i-1], nodes[i]
	}

	return &List{nodes[0], nodes[len(args)-1]}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	n := &Node{Value: v}

	if l.First() == nil {
		l.first, l.last = n, n
		return
	}

	l.first, l.first.prev, n.next = n, n, l.first
}

func (l *List) PushBack(v interface{}) {
	n := &Node{Value: v}

	if l.Last() == nil {
		l.first, l.last = n, n
		return
	}

	l.last, l.last.next, n.prev = n, n, l.last
}

func (l *List) PopFront() (result interface{}, err error) {
	switch l.First() {
	case nil:
		err = ErrEmptyList
	case l.Last():
		v := l.first.Value
		l.first, l.last = nil, nil
		result = v
	default:
		first := l.first
		l.first = first.next
		l.first.prev = nil
		result = first.Value
	}
	return result, err
}

func (l *List) PopBack() (result interface{}, err error) {
	switch l.Last() {
	case nil:
		err = ErrEmptyList
	case l.First():
		v := l.last.Value
		l.first, l.last = nil, nil
		result = v
	default:
		last := l.last
		l.last = last.prev
		l.last.next = nil
		result = last.Value
	}
	return result, err
}

func (l *List) Reverse() {
	n := l.First()

	if n == nil {
		return
	}

	l.first, l.last = l.last, l.first
	for {
		n.next, n.prev = n.prev, n.next
		if n.prev == nil {
			break
		}
		n = n.prev
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
