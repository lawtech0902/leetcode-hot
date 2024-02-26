/*
 * Author: robin-luo
 * Created time: 2024-02-22 15:24:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:27:10
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
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
