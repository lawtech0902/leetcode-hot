/*
__author__ = 'robin-luo'
__date__ = '2024/01/08 10:55'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree1(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	queue := []*TreeNode{p, q}
	for len(queue) != 0 {
		left := queue[0]
		queue = queue[1:]
		right := queue[0]
		queue = queue[1:]
		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		queue = append(queue, []*TreeNode{left.Left, right.Left, left.Right, right.Right}...)
	}

	return true
}
