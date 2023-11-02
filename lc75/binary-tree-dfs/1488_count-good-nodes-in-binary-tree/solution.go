/*
__author__ = 'robin-luo'
__date__ = '2023/11/02 11:05'
*/

package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, math.MinInt)
}

func dfs(root *TreeNode, pathMax int) int {
	if root == nil {
		return 0
	}

	res := 0
	if root.Val >= pathMax {
		res++
		pathMax = root.Val
	}

	res += dfs(root.Left, pathMax) + dfs(root.Right, pathMax)
	return res
}
