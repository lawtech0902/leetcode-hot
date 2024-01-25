/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 20:59'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

var front = &TreeNode{}

func flatten(root *TreeNode) {
	preOrder(root)
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}

	l := node.Left
	r := node.Right
	front.Right = node
	node.Left = nil
	front = node
	preOrder(l)
	preOrder(r)
}
