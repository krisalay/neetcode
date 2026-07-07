/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    var stringBuilder strings.Builder
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        stringVal := "#"
        if node != nil {
            stringVal = strconv.Itoa(node.Val)
            queue = append(queue, node.Left)
            queue = append(queue, node.Right) 
        }
        stringBuilder.WriteString(stringVal)
        stringBuilder.WriteString(",")
    }
    return stringBuilder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    parts := strings.Split(data, ",")
    idx := 0
    next := func() string {
        val := parts[idx]
        idx++
        return val
    }

    rootVal := next()
    if rootVal == "#" {
        return nil
    }

    root := &TreeNode{Val: mustAtoi(rootVal)}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        left := next()
        if left != "#" {
            node.Left = &TreeNode{Val: mustAtoi(left)}
            queue = append(queue, node.Left)
        }
        right := next()
        if right != "#" {
            node.Right = &TreeNode{Val: mustAtoi(right)}
            queue = append(queue, node.Right)
        }
    }
    return root
}

func mustAtoi(a string) int {
    n, _ := strconv.Atoi(a)
    return n
}
