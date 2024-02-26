/*
 * Author: robin-luo
 * Created time: 2024-02-22 19:54:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 19:56:47
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderSuccessor(root, p *TreeNode) *TreeNode {
	var successor *TreeNode
	if p.Right != nil {
		successor = p.Right
		for successor != nil {
			successor = successor.Left
		}

		return successor
	}

	node := root
	for node != nil {
		if node.Val > p.Val {
			successor = node
			node = node.Left
		} else {
			node = node.Right
		}
	}

	return successor
}
