/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:13:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 16:16:16
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	cur := root
	for cur != nil {
		if val > cur.Val {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				return root
			} else {
				cur = cur.Right
			}
		} else {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				return root
			} else {
				cur = cur.Left
			}
		}
	}

	return &TreeNode{Val: val}
}
