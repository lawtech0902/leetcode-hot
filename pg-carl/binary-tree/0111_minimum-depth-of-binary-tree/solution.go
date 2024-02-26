/*
 * Author: robin-luo
 * Created time: 2024-02-22 10:28:42
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 10:32:45
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	m1 := minDepth(root.Left)
	m2 := minDepth(root.Right)
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	} else {
		return min(m1, m2) + 1
	}
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		level++
		curLen := len(q)
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}

			if node.Left == nil && node.Right == nil {
				return level
			}
		}
	}

	return level
}
