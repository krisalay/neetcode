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

func largestRectangleArea(heights []int) int {
    prevSmallerIndex := getPrevSmallerIndex(heights)
    nextSmallerIndex := getNextSmallerIndex(heights)
	fmt.Println(prevSmallerIndex)
	fmt.Println(nextSmallerIndex)
    maxArea := 0
    for i := 0; i < len(heights); i++ {
        area := (nextSmallerIndex[i] - prevSmallerIndex[i] - 1) * heights[i]
        if area > maxArea {
            maxArea = area
        }
    }

    return maxArea
}

func getPrevSmallerIndex(arr []int) []int {
    result := make([]int, len(arr))
    var stack Stack
    for i := 0; i < len(arr); i++ {
        for !stack.IsEmpty() && arr[i] <= arr[stack.Top()] {
            stack.Pop()
        }
        if !stack.IsEmpty() {
            result[i] = stack.Top()
        } else {
            result[i] = -1
        }
        stack.Push(i)
    }
	return result
}

func getNextSmallerIndex(arr []int) []int {
    result := make([]int, len(arr))
    var stack Stack
    for i := len(arr)-1; i >= 0; i-- {
        for !stack.IsEmpty() && arr[i] <= arr[stack.Top()] {
            stack.Pop()
        }
        if !stack.IsEmpty() {
            result[i] = stack.Top()
        } else {
            result[i] = len(arr)
        }
        stack.Push(i)
    }
	return result
}
