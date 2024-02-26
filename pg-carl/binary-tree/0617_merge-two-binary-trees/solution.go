/*
 * Author: robin-luo
 * Created time: 2024-02-22 15:11:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:16:53
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}

	if root2 == nil {
		return root1
	}

	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}
