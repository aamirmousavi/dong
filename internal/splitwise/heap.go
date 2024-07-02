package splitwise

// maxHeap is a custom type for max heap using Node
type maxHeap[T comparable] []*node[T]

func (h maxHeap[T]) Len() int { return len(h) }
func (h maxHeap[T]) Less(i, j int) bool {
	return h[i].balance > h[j].balance
}
func (h maxHeap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap[T]) Push(x interface{}) {
	*h = append(*h, x.(*node[T]))
}

func (h *maxHeap[T]) Pop() interface{} {
	n := len(*h)
	result := (*h)[n-1]
	*h = (*h)[:n-1]
	return result
}
