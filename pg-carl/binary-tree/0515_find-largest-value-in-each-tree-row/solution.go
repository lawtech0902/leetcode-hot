/*
 * Author: robin-luo
 * Created time: 2024-02-22 10:02:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 10:05:17
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		maxVal := math.MinInt64
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			maxVal = max(maxVal, node.Val)
		}

		res = append(res, maxVal)
	}

	return res
}
