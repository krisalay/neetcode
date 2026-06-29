/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    var res int
    dfs(root, &res)
    return res
}

func dfs (curr *TreeNode, res *int) int {
    if curr == nil {
        return 0
    }
    left := dfs(curr.Left, res)
    right := dfs(curr.Right, res)
    *res = max(*res, left + right)
    return 1 + max(left, right)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
