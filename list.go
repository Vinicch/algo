package main

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (dll *DoublyLinkedList[T]) Prepend(item T) {
	dll.length++
	node := Node[T]{value: item}
	if dll.length == 1 {
		dll.head = &node
		dll.tail = &node
		return
	}

	node.next = dll.head
	dll.head.prev = &node
	dll.head = &node
}

func (dll *DoublyLinkedList[T]) InsertAt(item T, index int) {
	if index > dll.length {
		panic("the fuck, bro?")
	} else if index == dll.length {
		dll.Append(item)
		return
	} else if index == 0 {
		dll.Prepend(item)
		return
	}

	dll.length++
	curr := dll.head
	for i := 0; curr != nil && i < index; i++ {
		curr = curr.next
	}

	node := Node[T]{value: item}
	node.next = curr
	node.prev = curr.prev
	node.next.prev = &node
	node.prev.next = &node
}

func (dll *DoublyLinkedList[T]) Append(item T) {
	dll.length++
	node := Node[T]{value: item}
	if dll.length == 1 {
		dll.head = &node
		dll.tail = &node
		return
	}

	node.prev = dll.tail
	dll.tail.next = &node
	dll.tail = &node
}

func (dll *DoublyLinkedList[T]) Remove(item T) *T {
	curr := dll.head
	for i := 0; curr != nil && i < dll.length; i++ {
		if curr.value == item {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		return nil
	}

	dll.length--
	if dll.length == 0 {
		dll.head = nil
		dll.tail = nil
		return &curr.value
	}

	if curr == dll.head {
		dll.head = curr.next
	}

	if curr == dll.tail {
		dll.tail = curr.prev
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	curr.prev.next = curr.next
	curr.prev = nil
	curr.next = nil

	return &curr.value
}
