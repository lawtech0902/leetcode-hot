/*
__author__ = 'robin-luo'
__date__ = '2024/01/09 13:36'
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
	left, right := root.Left, root.Right
	root.Left = nil
	root.Right = left
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}

	cur.Right = right
}

var front = &TreeNode{}

func flatten1(root *TreeNode) {
	frontOrder(root)
}

func frontOrder(node *TreeNode) {
	if node == nil {
		return
	}

	left := node.Left
	right := node.Right
	front.Right = node
	node.Left = nil
	front = node
	frontOrder(left)
	frontOrder(right)
}
