package linkedlist

import (
	"errors"
)

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
	list := &List{}
	for _, n := range args {
		list.PushBack(n)
	}
	return list
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
		result = l.first.Value
		l.first, l.last = nil, nil
	default:
		result = l.first.Value
		l.first, l.first.next.prev = l.first.next, nil
	}
	return result, err
}

func (l *List) PopBack() (result interface{}, err error) {
	switch l.Last() {
	case nil:
		err = ErrEmptyList
	case l.First():
		result = l.last.Value
		l.first, l.last = nil, nil
	default:
		result = l.last.Value
		l.last, l.last.prev.next = l.last.prev, nil
	}
	return result, err
}

func (l *List) Reverse() {
	n := l.First()
	for n != nil {
		n.next, n.prev = n.prev, n.next
		n = n.prev
	}
	l.first, l.last = l.last, l.first
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
