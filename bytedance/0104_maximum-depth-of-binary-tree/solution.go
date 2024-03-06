/*
 * Author: robin-luo
 * Created time: 2024-03-05 11:11:41
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 11:24:22
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		curLen := len(q)
		level++
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return level
}
