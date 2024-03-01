/*
 * Author: robin-luo
 * Created time: 2024-02-29 17:12:14
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 17:34:19
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}

	prevSum = prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return prevSum
	}

	return dfs(root.Left, prevSum) + dfs(root.Right, prevSum)
}
