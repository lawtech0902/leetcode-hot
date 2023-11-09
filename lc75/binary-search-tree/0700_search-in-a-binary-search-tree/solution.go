/*
__author__ = 'robin-luo'
__date__ = '2023/11/08 16:59'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val < val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

func searchBST1(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if root.Val < val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}
