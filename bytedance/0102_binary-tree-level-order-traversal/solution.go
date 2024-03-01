/*
 * Author: robin-luo
 * Created time: 2024-02-29 14:59:54
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 15:02:39
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
		tempArr := make([]int, len(q))
		for i := range tempArr {
			node := q[0]
			q = q[1:]
			tempArr[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}
