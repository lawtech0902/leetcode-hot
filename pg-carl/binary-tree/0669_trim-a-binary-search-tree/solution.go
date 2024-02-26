/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:42:08
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 16:46:19
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func trimBST(root *TreeNode, low, high int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
