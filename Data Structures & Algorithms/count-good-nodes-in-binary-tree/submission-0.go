/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    count := 0
    countGoodNodes(root, math.MinInt, &count)
    return count
}

func countGoodNodes(root *TreeNode, maxValue int, count *int) {
    if root == nil {
        return
    }
    if root.Val >= maxValue {
        *count++
    }
    maxValue = max(maxValue, root.Val)
    countGoodNodes(root.Left, maxValue, count)
    countGoodNodes(root.Right, maxValue, count)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

