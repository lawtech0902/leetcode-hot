/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 20:59'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	leftNode := root.Left
	rightNode := root.Right
	root.Right = leftNode
	root.Left = nil

	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = rightNode
}
