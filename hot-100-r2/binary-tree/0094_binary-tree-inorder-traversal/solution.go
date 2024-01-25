/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 15:48'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// recursion
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}
