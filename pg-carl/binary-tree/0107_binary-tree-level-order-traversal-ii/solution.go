/*
 * Author: robin-luo
 * Created time: 2024-02-21 20:57:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 21:05:00
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
