type State struct {
    Val int
    Min int
}

type MinStack struct {
    store []State
}

func Constructor() MinStack {
    return MinStack{
        store: []State{},
    }
}

func (this *MinStack) Push(val int) {
    if len(this.store) == 0 {
        this.store = append(this.store, State{Val: val, Min: val})
        return
    }
    topState := this.top()
    minVal := min(topState.Min, val)
    this.store = append(this.store, State{Val: val, Min: minVal})
}

func (this *MinStack) Pop() {
    this.store = this.store[:len(this.store)-1]
}

func (this *MinStack) top() State {
    return this.store[len(this.store)-1]
}

func (this *MinStack) Top() int {
    state := this.top()
    return state.Val
}

func (this *MinStack) GetMin() int {
    topState := this.top()
    return topState.Min
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
