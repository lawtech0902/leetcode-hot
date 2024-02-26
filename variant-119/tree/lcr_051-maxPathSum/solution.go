/*
 * Author: robin-luo
 * Created time: 2024-02-22 17:13:25
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 17:17:28
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)
		maxSum = max(maxSum, leftGain+root.Val+rightGain)
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}
