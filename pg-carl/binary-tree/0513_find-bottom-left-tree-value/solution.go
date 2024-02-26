/*
 * Author: robin-luo
 * Created time: 2024-02-22 13:53:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 13:55:42
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	var res int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			if i == 0 {
				res = node.Val
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
