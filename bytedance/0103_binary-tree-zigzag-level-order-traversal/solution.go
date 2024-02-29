/*
 * Author: robin-luo
 * Created time: 2024-02-29 10:20:04
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 10:24:18
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	level := 0
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

		if level%2 == 1 {
			for i, j := 0, len(tempArr)-1; i < j; i, j = i+1, j-1 {
				tempArr[i], tempArr[j] = tempArr[j], tempArr[i]
			}
		}

		res = append(res, tempArr)
		level++
	}

	return res
}
