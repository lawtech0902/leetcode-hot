/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:17:56
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 16:26:21
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Right != nil && root.Right.Left != nil {
			node := root.Right
			for node.Left != nil {
				node = node.Left
			}

			node.Left = root.Left
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		if root.Right != nil {
			root.Right.Left = root.Left
			return root.Right
		}
	}

	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}

	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}
