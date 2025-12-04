package list

import "fmt"

type DoublyLinkedListNode[T any] struct {
	Value T
	Prev  *DoublyLinkedListNode[T]
	Next  *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T any] struct {
	Head       *DoublyLinkedListNode[T]
	TotalItems int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	ls := DoublyLinkedList[T]{}
	return &ls
}

func (l *DoublyLinkedList[T]) Add(pos int, value T) error {
	if pos > l.TotalItems {
		return ErrInvalidPosition
	}

	newNode := &DoublyLinkedListNode[T]{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
	if pos == 0 {
		newNode.Next = l.Head
		l.Head = newNode
	} else if l.TotalItems == pos {
		previousNode := l.nodeAt(pos - 1)
		nextNode := previousNode.Next
		previousNode.Next = newNode
		newNode.Prev = previousNode
		newNode.Next = nextNode
	} else {
		currentNode := l.nodeAt(pos)
		previousNode := currentNode.Prev
		previousNode.Next = newNode
		newNode.Prev = previousNode
		newNode.Next = currentNode
		currentNode.Prev = newNode
	}

	l.TotalItems++
	return nil
}

func (l *DoublyLinkedList[T]) Delete(pos int) (T, error) {
	var item T
	if pos >= l.TotalItems {
		return item, ErrInvalidPosition
	}

	if l.TotalItems == 0 {
		return item, ErrInvalidPosition
	}

	if pos == 0 {
		head := l.Head
		nextNode := l.Head.Next
		l.Head = nextNode
		nextNode.Prev = l.Head
		head.Next = nil
		head.Prev = nil
		item = head.Value
	} else {
		currentNode := l.nodeAt(pos)
		previousNode := currentNode.Prev
		nextNode := currentNode.Next

		previousNode.Next = nextNode

		// Last node will have nil next node. but middle nodes will have.
		if pos < l.TotalItems-1 {
			nextNode.Prev = currentNode
		}

		currentNode.Next = nil
		currentNode.Prev = nil
		item = currentNode.Value
	}
	l.TotalItems--
	return item, nil
}

func (l *DoublyLinkedList[T]) Value(pos int) (T, error) {
	var t T
	if pos >= l.TotalItems {
		return t, ErrInvalidPosition
	}
	n := l.nodeAt(pos)
	return n.Value, nil
}

func (l *DoublyLinkedList[T]) nodeAt(pos int) *DoublyLinkedListNode[T] {

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

func (l *DoublyLinkedList[T]) Len() int {
	return l.TotalItems
}

func (l *DoublyLinkedList[T]) String() string {
	s := ""
	h := l.Head
	for h != nil {
		s += fmt.Sprintf("%v->", h.Value)
		h = h.Next
	}
	s += "NIL"
	return s
}
