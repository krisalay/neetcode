type Entity struct {
	Point []int
	Distance float64
}

type MaxHeap []Entity
func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Distance > h[j].Distance
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(Entity))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	val := old[n-1]
	*h = old[:n-1]
	return val
}

func kClosest(points [][]int, k int) [][]int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, point := range points {
		distance := getDistanceFromOrigin(point)
		heap.Push(maxHeap, Entity{Point: point, Distance: distance})
		if len(*maxHeap) > k {
			heap.Pop(maxHeap)
		}
	}
	result := make([][]int, k)
	ptr := 0
	for len(*maxHeap) > 0 {
		entity := (heap.Pop(maxHeap)).(Entity)
		result[ptr] = entity.Point
		ptr++
	}
	return result
}

func getDistanceFromOrigin(point []int) float64 {
	return math.Sqrt(float64(point[0]*point[0]) + float64(point[1]*point[1]))
}
