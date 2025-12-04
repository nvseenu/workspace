package list

import (
	"fmt"
)

type LinkedListNode[T any] struct {
	Value T
	Next  *LinkedListNode[T]
}

type LinkedList[T any] struct {
	Head          *LinkedListNode[T]
	TotalElements int
}

func NewLinkedList[T any]() *LinkedList[T] {
	ls := LinkedList[T]{}
	return &ls
}

func (l *LinkedList[T]) Add(pos int, value T) error {
	if pos > l.TotalElements {
		return ErrInvalidPosition
	}

	n := &LinkedListNode[T]{
		Value: value,
		Next:  nil,
	}
	if pos == 0 {
		n.Next = l.Head
		l.Head = n
	} else {
		previousNode := l.nodeAt(pos - 1)
		nextNode := previousNode.Next
		previousNode.Next = n
		n.Next = nextNode
	}
	l.TotalElements++
	return nil
}

func (l *LinkedList[T]) Delete(pos int) (T, error) {
	var elm T
	if pos > l.TotalElements {
		return elm, ErrInvalidPosition
	}

	if l.TotalElements == 0 {
		return elm, ErrInvalidPosition
	}

	if pos == 0 {
		head := l.Head
		nextNode := l.Head.Next
		l.Head = nextNode
		elm = head.Value
	} else {
		previousNode := l.nodeAt(pos - 1)
		currentNode := previousNode.Next
		nextNode := currentNode.Next
		previousNode.Next = nextNode
		currentNode.Next = nil
		elm = currentNode.Value
	}
	l.TotalElements--
	return elm, nil
}

func (l *LinkedList[T]) Value(pos int) (T, error) {
	var t T
	if pos >= l.TotalElements {
		return t, ErrInvalidPosition
	}
	n := l.nodeAt(pos)
	return n.Value, nil
}

func (l *LinkedList[T]) nodeAt(pos int) *LinkedListNode[T] {

	head := l.Head
	count := 0

	for head != nil {
		if count == pos {
			return head
		}
		head = head.Next
		count++
	}
	return nil
}

func (l *LinkedList[T]) Len() int {
	return l.TotalElements
}

func (l *LinkedList[T]) String() string {
	s := ""
	h := l.Head
	for h != nil {
		s += fmt.Sprintf("%v->", h.Value)
		h = h.Next
	}
	s += "NIL"
	return s
}
