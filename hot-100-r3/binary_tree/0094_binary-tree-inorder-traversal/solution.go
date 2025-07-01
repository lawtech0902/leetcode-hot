/*
__author__ = 'robin-luo'
__date__ = '2025/06/30 14:24'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}

	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var (
		res       []int
		traversal func(root *TreeNode)
	)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}

	traversal(root)
	return res
}
