/*
 * Author: robin-luo
 * Created time: 2024-03-05 14:27:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 14:54:55
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

func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	return root
}
