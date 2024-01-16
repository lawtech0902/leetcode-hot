/*
__author__ = 'robin-luo'
__date__ = '2024/01/12 16:12'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// bfs
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		curLen := len(q)
		var levelSum float64
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			levelSum += float64(node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, levelSum/float64(curLen))
	}

	return res
}
