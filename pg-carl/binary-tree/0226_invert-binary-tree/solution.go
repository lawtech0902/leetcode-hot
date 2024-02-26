/*
 * Author: robin-luo
 * Created time: 2024-02-22 10:39:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 10:40:17
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
