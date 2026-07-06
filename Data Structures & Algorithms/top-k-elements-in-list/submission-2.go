type Entity struct {
	Val int
	Freq int
}
type MinHeap []Entity

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Freq < h[j].Freq
}

func (h *MinHeap) Push(x any) {
    *h = append(*h, x.(Entity))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	pq := &MinHeap{}
	heap.Init(pq)

	for val, freq := range freqMap {
		heap.Push(pq, Entity{Val: val, Freq: freq})

		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	ans := []int{}
	for pq.Len() > 0 {
		e := heap.Pop(pq).(Entity)
		ans = append(ans, e.Val)
	}
	return ans
}
