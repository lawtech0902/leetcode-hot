/*
__author__ = 'robin-luo'
__date__ = '2025/07/11 10:03'
*/

package solution

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var (
		res       []int
		traversal func(node *TreeNode)
	)

	traversal = func(node *TreeNode) {
		if root == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}

	traversal(root)
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	slices.Reverse(res)
	return res
}
