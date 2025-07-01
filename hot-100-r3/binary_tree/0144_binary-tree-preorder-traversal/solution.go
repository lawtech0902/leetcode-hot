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
		traversal func(root *TreeNode)
	)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		traversal(root.Left)
		traversal(root.Right)
	}

	traversal(root)
	return res
}
