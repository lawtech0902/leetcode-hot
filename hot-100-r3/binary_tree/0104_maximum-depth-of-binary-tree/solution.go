/*
__author__ = 'robin-luo'
__date__ = '2025/07/14 09:58'
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

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
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
