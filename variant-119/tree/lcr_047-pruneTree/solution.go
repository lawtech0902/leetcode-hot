/*
 * Author: robin-luo
 * Created time: 2024-02-21 15:43:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 15:56:33
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
