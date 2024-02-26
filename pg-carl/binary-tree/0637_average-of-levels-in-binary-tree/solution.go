/*
 * Author: robin-luo
 * Created time: 2024-02-22 09:48:00
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 09:48:25
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)
		var levelSum float64
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			levelSum += float64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, levelSum/float64(size))
	}

	return res
}
