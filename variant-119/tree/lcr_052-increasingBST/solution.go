/*
 * Author: robin-luo
 * Created time: 2024-02-22 17:28:51
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 19:52:43
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	cur := dummy
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		// 在中序遍历的过程中修改节点指向
		cur.Right = node
		node.Left = nil
		cur = node

		inorder(node.Right)
	}

	inorder(root)
	return dummy.Right
}
