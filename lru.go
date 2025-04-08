package main

type LRUNode[K comparable, V any] struct {
	key   K
	value V
	next  *LRUNode[K, V]
	prev  *LRUNode[K, V]
}

type LRU[K comparable, V any] struct {
	length   int
	capacity int
	head     *LRUNode[K, V]
	tail     *LRUNode[K, V]
	lookup   map[K]*LRUNode[K, V]
}

func NewLRU[K comparable, V any](capacity int) LRU[K, V] {
	return LRU[K, V]{lookup: map[K]*LRUNode[K, V]{}, capacity: capacity}
}

func (lru *LRU[K, V]) Update(key K, value V) {
	node, ok := lru.lookup[key]
	if ok {
		lru.detach(node)
		lru.prepend(node)
		node.value = value
		return
	}

	node = &LRUNode[K, V]{key: key, value: value}
	lru.lookup[key] = node
	lru.prepend(node)
	lru.length++
	lru.evict()
}

func (lru *LRU[K, V]) Get(key K) *V {
	node, ok := lru.lookup[key]
	if !ok {
		return nil
	}

	lru.detach(node)
	lru.prepend(node)
	return &node.value
}

func (lru *LRU[K, V]) detach(node *LRUNode[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if node == lru.head {
		lru.head = lru.head.next
	}

	if node == lru.tail {
		lru.tail = lru.tail.prev
	}

	node.prev = nil
	node.next = nil
}

func (lru *LRU[K, V]) prepend(node *LRUNode[K, V]) {
	if lru.head == nil {
		lru.head = node
		lru.tail = node
		return
	}

	lru.head.prev = node
	node.next = lru.head
	lru.head = node
}

func (lru *LRU[K, V]) evict() {
	if lru.length <= lru.capacity {
		return
	}

	delete(lru.lookup, lru.tail.key)
	lru.detach(lru.tail)
	lru.length--
}
