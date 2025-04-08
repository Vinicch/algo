package main

type MinHeap struct {
	length int
	data   []int
}

func (h *MinHeap) Insert(value int) {
	h.data[h.length] = value
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	h.length--
	head := h.data[0]
	if h.length == 0 {
		h.data = []int{}
		return head
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)
	return head
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	parent := h.data[p]

	if parent > h.data[idx] {
		h.data[p] = h.data[idx]
		h.data[idx] = parent
		h.heapifyUp(p)
	}
}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= h.length {
		return
	}

	l := leftChild(idx)
	r := rightChild(idx)

	if l >= h.length {
		return
	}

	left, right, curr := h.data[l], h.data[r], h.data[idx]
	if left > right && curr > right {
		h.data[r] = curr
		h.data[idx] = right
		h.heapifyDown(r)
	} else if right > left && curr > left {
		h.data[l] = curr
		h.data[idx] = left
		h.heapifyDown(l)
	}
}

func parent(index int) int {
	return (index - 1) / 2
}

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}
