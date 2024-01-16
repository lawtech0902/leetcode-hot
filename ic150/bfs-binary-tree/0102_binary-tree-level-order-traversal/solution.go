/*
__author__ = 'robin-luo'
__date__ = '2024/01/15 11:39'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// bfs
func levelOrder1(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tempArr := make([]int, len(q))
		for i := range tempArr {
			node := q[0]
			q = q[1:]
			tempArr[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}

// bfs
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	return inner(root, 0, res)
}

func inner(root *TreeNode, level int, res [][]int) [][]int {
	if root == nil {
		return res
	}

	if level >= len(res) {
		res = append(res, []int{})
	}

	if root != nil {
		res[level] = append(res[level], root.Val)
	}

	inner(root.Left, level+1, res)
	inner(root.Right, level+1, res)
	return res
}
