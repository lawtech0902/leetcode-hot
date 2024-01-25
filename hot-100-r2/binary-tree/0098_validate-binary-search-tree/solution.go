/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 20:41'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, -1<<63, 1<<63-1)
}

func validate(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if root.Val <= left || root.Val >= right {
		return false
	}

	return validate(root.Left, left, root.Val) && validate(root.Right, root.Val, right)
}
