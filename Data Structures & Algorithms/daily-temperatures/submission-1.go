type Stack []int

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}

func (s *Stack) Push(val int) {
    *s = append(*s, val)
}

func (s *Stack) Pop() {
    *s = (*s)[:len(*s)-1]
}

func (s *Stack) Top() int {
	return (*s)[len(*s)-1]
}

func dailyTemperatures(temperatures []int) []int {
    ngi := getNextGreaterIndex(temperatures)
    result := make([]int, len(temperatures))
    for i, idx := range ngi {
        result[i] = idx - i
    }
    return result
}

func getNextGreaterIndex(arr []int) []int {
    result := make([]int, len(arr))
    var stack Stack
    for i := len(arr)-1; i >= 0; i-- {
        for !stack.IsEmpty() && arr[i] >= arr[stack.Top()] {
            stack.Pop()
        }
        if !stack.IsEmpty() {
            result[i] = stack.Top()
        } else {
            result[i] = i
        }
        stack.Push(i)
    }
    return result
}
