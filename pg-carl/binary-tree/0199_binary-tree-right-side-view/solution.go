/*
 * Author: robin-luo
 * Created time: 2024-02-22 09:32:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 09:39:30
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
		for curLen != 0 {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			curLen--
			if curLen == 0 {
				res = append(res, node.Val)
			}
		}
	}

	return res
}
