/*
__author__ = 'robin-luo'
__date__ = '2024/01/12 14:45'
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
	for len(q) != 0 {
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
