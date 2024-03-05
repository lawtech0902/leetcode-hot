/*
 * Author: robin-luo
 * Created time: 2024-03-04 14:14:00
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 14:25:06
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxVal := 0
	dfs(root, &maxVal)
	return maxVal
}

func dfs(node *TreeNode, maxVal *int) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left, maxVal)
	r := dfs(node.Right, maxVal)
	l1, r1 := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		l1 = l + 1
	}

	if node.Right != nil && node.Right.Val == node.Val {
		r1 = r + 1
	}

	*maxVal = max(*maxVal, l1+r1)
	return max(l1, r1)
}
