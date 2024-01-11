/*
__author__ = 'robin-luo'
__date__ = '2024/01/08 15:11'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return helper(root.Left, root.Right)
}

func helper(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return helper(p.Left, q.Right) && helper(p.Right, q.Left)
}
