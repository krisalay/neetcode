type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)

	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}

	for len(*maxHeap) > 1 {
		firstWeight := (heap.Pop(maxHeap)).(int)
		secondWeight := (heap.Pop(maxHeap)).(int)

		if firstWeight != secondWeight {
			newWeight := abs(firstWeight - secondWeight)
			heap.Push(maxHeap, newWeight)
		}
	}

	if len(*maxHeap) == 0 {
		return 0
	}
	return (heap.Pop(maxHeap)).(int)
}

func abs(a int) int {
	if a < 0 {
		return -1*a
	}
	return a
}
