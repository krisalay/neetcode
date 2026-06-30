/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    result := math.MinInt
    calculateMaxPathSum(root, &result)
    return result
}

func calculateMaxPathSum(root *TreeNode, result *int) int {
    if root == nil {
        return 0
    }
    leftPathSum := calculateMaxPathSum(root.Left, result)
    rightPathSum := calculateMaxPathSum(root.Right, result)
	if leftPathSum < 0 {
		leftPathSum = 0
	}
	if rightPathSum < 0 {
		rightPathSum = 0
	}
    *result = max(*result, root.Val + leftPathSum + rightPathSum)
    return max(leftPathSum, rightPathSum) + root.Val
}
