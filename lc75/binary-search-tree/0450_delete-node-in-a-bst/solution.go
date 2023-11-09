/*
__author__ = 'robin-luo'
__date__ = '2023/11/09 09:28'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		min := getMin(root.Right)
		root.Val = min.Val
		root.Right = deleteNode(root.Right, min.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	for root.Left != nil {
		root = root.Left
	}

	return root
}
