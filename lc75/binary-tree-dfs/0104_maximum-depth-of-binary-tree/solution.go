/*
__author__ = 'robin-luo'
__date__ = '2023/11/01 14:13'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func maxDepth1(root *TreeNode) int {

}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
