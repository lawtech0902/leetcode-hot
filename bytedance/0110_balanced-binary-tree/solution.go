/*
 * Author: robin-luo
 * Created time: 2024-03-05 11:06:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 11:10:52
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getDepth(root.Left)
	r := getDepth(root.Right)
	if abs(l-r) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(getDepth(node.Left), getDepth(node.Right))
}

func abs(a int) int {
	if a >= 0 {
		return a
	}

	return -a
}
