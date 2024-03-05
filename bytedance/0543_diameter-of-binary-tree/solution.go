/*
 * Author: robin-luo
 * Created time: 2024-03-04 13:51:25
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 14:01:21
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	dfs(root, &d)
	return d
}

func dfs(node *TreeNode, d *int) int {
	if node == nil {
		return 0
	}

	l := dfs(node.Left, d)
	r := dfs(node.Right, d)
	*d = max(*d, l+r)
	return max(l, r) + 1
}
