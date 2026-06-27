func dailyTemperatures(temperatures []int) []int {
    out := make([]int, len(temperatures))
    stack := NewStack()
    
    for i, temp := range temperatures {
        for len(stack.store) > 0 {
            top := stack.Top()
            if temp > top.temp {
                out[top.idx] = i-top.idx
                stack.Pop()
                continue
            }
            break
        }
        stack.Push(temp, i)
    }
    return out
}

type State struct {
    temp int
    idx int
}

type Stack struct {
    store []State
}

func NewStack() *Stack {
    return &Stack{
        store: []State{},
    }
}

func (s *Stack) Top()  State {
    return s.store[len(s.store)-1]
}

func (s *Stack) Pop() {
    s.store = s.store[:len(s.store)-1]
}

func (s *Stack) Push(temp, idx int) {
    s.store = append(s.store, State{temp: temp, idx: idx})
}
