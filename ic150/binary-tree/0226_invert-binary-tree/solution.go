/*
__author__ = 'robin-luo'
__date__ = '2024/01/08 14:59'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}
