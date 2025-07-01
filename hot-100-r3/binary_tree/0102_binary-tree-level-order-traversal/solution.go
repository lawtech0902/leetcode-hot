/*
__author__ = 'robin-luo'
__date__ = '2025/07/01 10:06'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := make([]int, len(q))
		for i := range temp {
			node := q[0]
			q = q[1:]
			temp[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, temp)
	}

	return res
}
