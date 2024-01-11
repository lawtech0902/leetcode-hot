/*
__author__ = 'robin-luo'
__date__ = '2024/01/10 9:43'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, preSum int) int
	dfs = func(root *TreeNode, preSum int) int {
		if root == nil {
			return 0
		}

		preSum = preSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return preSum
		}

		return dfs(root.Left, preSum) + dfs(root.Right, preSum)
	}

	return dfs(root, 0)
}
