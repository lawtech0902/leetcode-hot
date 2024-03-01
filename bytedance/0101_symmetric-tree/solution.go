/*
 * Author: robin-luo
 * Created time: 2024-02-29 17:36:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 17:37:45
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return helper(l.Left, r.Right) && helper(l.Right, r.Left)
}
