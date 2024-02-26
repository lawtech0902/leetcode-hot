/*
 * Author: robin-luo
 * Created time: 2024-02-22 11:12:05
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 11:16:33
 */

package solution

import "math"

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
	if math.Abs(float64(l-r)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getDepth(root.Left), getDepth(root.Right))
}
