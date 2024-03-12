/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 11:16'
*/

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxVal := math.MinInt64
	dfs(root, &maxVal)
	return maxVal
}

func dfs(node *TreeNode, maxVal *int) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left, maxVal)
	r := dfs(node.Right, maxVal)
	*maxVal = max(*maxVal, node.Val+l+r)
	return max(0, max(l, r)+node.Val)
}
