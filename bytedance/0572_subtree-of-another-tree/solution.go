/*
 * Author: robin-luo
 * Created time: 2024-02-28 15:07:45
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 15:08:25
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSubtree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	if s.Val == t.Val && isSame(s, t) {
		return true
	}

	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
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
