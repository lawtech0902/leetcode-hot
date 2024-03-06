/*
 * Author: robin-luo
 * Created time: 2024-03-05 09:53:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 10:00:34
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isCompleteTree(root *TreeNode) bool {
	// 标记层序遍历的过程中是否有遇到空节点
	empty := false
	q := []*TreeNode{root}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			// 遇到空节点，把标记改成true
			empty = true
		} else {
			// 此时是遍历非空节点，在非空节点之前出现了空节点，就说明并不是完全二叉树
			if empty {
				return false
			}

			// 添加左右子节点，无论是否为空
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}

	return true
}
