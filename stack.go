package main

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func (s *Stack[T]) Push(value T) {
	s.length++
	node := Node[T]{value: value}
	if s.length == 1 {
		s.head = &node
		return
	}

	node.next = s.head
	s.head = &node
}

func (s *Stack[T]) Pop() *T {
	if s.head == nil {
		return nil
	}

	h := s.head
	s.head = s.head.next
	h.next = nil
	s.length--

	return &h.value
}

func (s *Stack[T]) Peek() *T {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}
