type PriorityQueue []int
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i],pq[j] = pq[j],pq[i]
}

func (pq *PriorityQueue) Push(val any) {
	*pq = append(*pq, val.(int))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
    val := old[n-1]
    *pq = old[:n-1]
    return val
}

type KthLargest struct {
    capacity int
	pq *PriorityQueue
}


func Constructor(k int, nums []int) KthLargest {
    pq := &PriorityQueue{}
	heap.Init(pq)
	for _, num := range nums {
		heap.Push(pq, num)
		if len(*pq) > k {
			heap.Pop(pq)
		}
	}
	return KthLargest{
		capacity: k,
		pq: pq,
	}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.pq, val)
    if len(*this.pq) > this.capacity {
        heap.Pop(this.pq)
    }
    return (*this.pq)[0]  // min of heap = kth largest
}
