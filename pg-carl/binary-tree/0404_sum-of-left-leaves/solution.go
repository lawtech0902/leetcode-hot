/*
 * Author: robin-luo
 * Created time: 2024-02-22 11:29:17
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 11:30:39
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
}
