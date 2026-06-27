func evalRPN(tokens []string) int {
    stack := []int{}
    for _, token := range tokens {
        d, ok := isNumber(token)
        if ok {
            stack = append(stack, d)
            continue
        }
        secondElem := stack[len(stack)-1]
        firstElem := stack[len(stack)-2]
        stack = stack[:len(stack)-2]
        res := firstElem + secondElem
        if token == "-" {
            res = firstElem - secondElem
        }
        if token == "*" {
            res = firstElem * secondElem
        }
        if token == "/" {
            res = firstElem / secondElem
        }
        stack = append(stack, res)
    }
    return stack[0]
}

func isNumber(s string) (int, bool) {
    d, err := strconv.Atoi(s)
    return d, err == nil
}
