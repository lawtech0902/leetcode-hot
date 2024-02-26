/*
 * Author: robin-luo
 * Created time: 2024-02-22 15:17:30
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:19:02
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}
