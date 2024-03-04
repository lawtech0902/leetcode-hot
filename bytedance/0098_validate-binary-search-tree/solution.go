/*
 * Author: robin-luo
 * Created time: 2024-03-01 17:06:33
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 17:10:13
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

func validate(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}

	if root.Val <= l || root.Val >= r {
		return false
	}

	return validate(root.Left, l, root.Val) && validate(root.Right, root.Val, r)
}
