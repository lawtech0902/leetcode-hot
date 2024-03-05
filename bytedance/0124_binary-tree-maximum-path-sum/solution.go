/*
 * Author: robin-luo
 * Created time: 2024-02-29 15:41:32
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 14:13:08
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
