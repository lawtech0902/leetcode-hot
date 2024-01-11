/*
__author__ = 'robin-luo'
__date__ = '2024/01/10 13:57'
*/

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	dfs(root, &maxSum)
	return maxSum
}

func dfs(root *TreeNode, maxSum *int) int {
	if root == nil {
		return 0
	}

	l := dfs(root.Left, maxSum)
	r := dfs(root.Right, maxSum)
	*maxSum = max(*maxSum, l+r+root.Val, root.Val+l, root.Val+r, root.Val)
	return max(root.Val, root.Val+l, root.Val+r)
}
