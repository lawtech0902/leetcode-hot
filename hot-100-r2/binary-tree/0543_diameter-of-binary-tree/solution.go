/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 17:18'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	d := 0
	depth(root, &d)
	return d
}

func depth(node *TreeNode, d *int) int {
	if node == nil {
		return 0
	}

	l := depth(node.Left, d)
	r := depth(node.Right, d)
	*d = max(*d, l+r)
	return max(l, r) + 1
}
