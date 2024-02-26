/*
 * Author: robin-luo
 * Created time: 2024-02-22 13:58:04
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 14:01:50
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
