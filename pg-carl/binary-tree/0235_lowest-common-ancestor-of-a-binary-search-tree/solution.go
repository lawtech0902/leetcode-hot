/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:01:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 16:07:48
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}

	return cur
}
