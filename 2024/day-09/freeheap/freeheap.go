package freeheap

import "container/heap"

type FreeHeap []int

// interface sort.Interface
func (h FreeHeap) Len() int {
	return len(h)
}

func (h FreeHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h FreeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// interface heap.Interface
func (h *FreeHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *FreeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var _ heap.Interface = (*FreeHeap)(nil)
