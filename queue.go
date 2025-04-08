package main

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.length++
	node := Node[T]{value: value}
	if q.length == 1 {
		q.head = &node
		q.tail = &node
		return
	}

	q.tail.next = &node
	q.tail = &node
}

func (q *Queue[T]) Dequeue() *T {
	if q.head == nil {
		return nil
	}

	h := q.head
	q.head = q.head.next
	h.next = nil
	q.length--
	if q.length == 0 {
		q.tail = nil
	}

	return &h.value
}

func (q *Queue[T]) Peek() *T {
	if q.head == nil {
		return nil
	}
	return &q.head.value
}
