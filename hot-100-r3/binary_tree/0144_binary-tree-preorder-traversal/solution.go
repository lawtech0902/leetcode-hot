/*
__author__ = 'robin-luo'
__date__ = '2025/07/01 10:14'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var (
		res       []int
		traversal func(node *TreeNode)
	)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

func preorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}
