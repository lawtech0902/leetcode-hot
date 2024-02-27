/*
 * Author: robin-luo
 * Created time: 2024-02-26 20:27:30
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-26 20:31:10
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSubTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && isSame(s, t) {
		return true
	}

	return isSubTree(s.Left, t) || isSubTree(s.Right, t)
}

func isSame(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil || s.Val != t.Val {
		return false
	}

	return isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}
