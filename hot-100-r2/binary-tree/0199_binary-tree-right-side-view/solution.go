/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 20:54'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLen := len(queue)
		for curLen != 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			curLen--
			if curLen == 0 {
				res = append(res, node.Val)
			}
		}
	}

	return res
}
