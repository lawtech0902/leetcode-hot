/*
 * Author: robin-luo
 * Created time: 2024-02-22 19:57:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 19:59:24
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	pre := 0
	traversal(root, &pre)
	return root
}

func traversal(root *TreeNode, pre *int) {
	if root == nil {
		return
	}

	traversal(root.Right, pre)
	root.Val += *pre
	*pre = root.Val
	traversal(root.Left, pre)
}
