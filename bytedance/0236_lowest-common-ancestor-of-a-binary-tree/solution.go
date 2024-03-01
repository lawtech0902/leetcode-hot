/*
 * Author: robin-luo
 * Created time: 2024-02-29 10:57:39
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 10:59:04
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	} else {
		return r
	}
}
