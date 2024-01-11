/*
__author__ = 'robin-luo'
__date__ = '2024/01/08 10:39'
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
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		level++
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return level
}
