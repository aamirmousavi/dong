package core

// MaxHeap is a custom type for max heap using Node
type MaxHeap []*Node

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].FinalBalance > h[j].FinalBalance
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	result := (*h)[n-1]
	*h = (*h)[:n-1]
	return result
}
