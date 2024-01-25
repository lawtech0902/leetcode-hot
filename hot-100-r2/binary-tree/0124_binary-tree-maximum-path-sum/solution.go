/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 11:16'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 63
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		maxSum = max(maxSum, leftGain+node.Val+rightGain)
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
