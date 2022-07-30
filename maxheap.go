package gostruct

// Heap is a binary max heap in generic way
// comp is a function that compares two values and returns true if a is smaller than b
type Heap[T any] struct {
	data []T
	comp func(a, b T) bool
}

func NewHeap[T any](comp func(a, b T) bool) *Heap[T] {
	return &Heap[T]{
		data: make([]T, 0),
		comp: comp,
	}
}

// Push an value into the already sorted heap
func (h *Heap[T]) Push(v T) {
	h.data = append(h.data, v)
	h.heapifyUp((len(h.data) - 1))
}

// heapifyUp moves the value at index i up to its correct position in the heap
// Assume other values are already in the correct position in the heap except for i
func (h *Heap[T]) heapifyUp(i int) {
	for h.comp(h.data[parentIndex(i)], h.data[i]) {
		h.swap(i, parentIndex(i))
		i = parentIndex(i)
	}
}

// Pop return the top value of the heap and rearrange the heap
// If the heap is empty, return zero value and false
func (h *Heap[T]) Pop() (T, bool) {
	var val T
	if h.Size() == 0 {
		return val, false
	}
	val = h.data[0]
	h.swap(0, h.Size()-1)
	h.data = h.data[:h.Size()-1]
	h.heapifyDown(0)

	return val, true
}

// heapifyDown moves the value at index i down to its correct position in the heap
// Assume the left and right branch are already in the correct position in the heap except for i
//  assumes that the binary trees rooted at LEFT.i/ and RIGHT.i/ are max- heaps, but that AŒi􏰌 might be smaller than its children, thus violating the max-heap property
func (h *Heap[T]) heapifyDown(i int) {
	l, r := leftChildIndex(i), rightChildIndex(i)
	largest := i
	if l < len(h.data) && h.comp(h.data[i], h.data[l]) {
		largest = l
	}
	if r < len(h.data) && h.comp(h.data[largest], h.data[r]) {
		largest = r
	}
	if largest != i {
		h.swap(i, largest)
		h.heapifyDown(largest)
	}
}

// swap swaps the values at indices i and j
func (h *Heap[T]) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// leftChildIndex returns the index of left child of i
func leftChildIndex(i int) int {
	return 2*i + 1
}

// rightChildIndex returns the index of right child of i
func rightChildIndex(i int) int {
	return 2*i + 2
}

// parentIndex returns the index of parent of i
func parentIndex(i int) int {
	return (i - 1) / 2
}

// Peek returns the top value of the heap without rearrange the heap
func (h *Heap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		var val T
		return val, false
	}
	return h.data[0], true
}

func (h *Heap[T]) Size() int {
	return len(h.data)
}
