/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return validateBST(root, math.MinInt, math.MaxInt)
}

func validateBST(root *TreeNode, minVal, maxVal int) bool {
    if root == nil {
        return true
    }

    if root.Val <= minVal || root.Val >= maxVal {
        return false
    }

    return validateBST(root.Left, minVal, root.Val) && validateBST(root.Right, root.Val, maxVal)
}
