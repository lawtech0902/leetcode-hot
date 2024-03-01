/*
 * Author: robin-luo
 * Created time: 2024-02-29 14:45:56
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 14:48:51
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		for curLen > 0 {
			node := q[0]
			q = q[1:]
			curLen--
			if curLen == 0 {
				res = append(res, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}
