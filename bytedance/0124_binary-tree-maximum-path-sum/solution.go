/*
 * Author: robin-luo
 * Created time: 2024-02-29 15:41:32
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 16:02:53
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	var maxGain func(root *TreeNode) int
	maxGain = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftGain := max(maxGain(root.Left), 0)
		rightGain := max(maxGain(root.Right), 0)
		maxSum = max(maxSum, leftGain+root.Val+rightGain)
		return root.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

func maxPathSum1(root *TreeNode) int {
	res := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, node.Val+left+right)
		return max(0, max(left, right)+node.Val)
	}

	dfs(root)
	return res
}
