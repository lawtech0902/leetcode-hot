/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:59:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 17:00:58
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	convert(root, &sum)
	return root
}

func convert(node *TreeNode, sum *int) {
	if node == nil {
		return
	}

	convert(node.Right, sum)
	node.Val += *sum
	*sum = node.Val
	convert(node.Left, sum)
}
