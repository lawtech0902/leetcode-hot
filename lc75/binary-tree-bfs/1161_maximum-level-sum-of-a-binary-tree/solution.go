/*
__author__ = 'robin-luo'
__date__ = '2023/11/06 11:01'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func maxLevelSum(root *TreeNode) int {
	res, maxSum := 1, root.Val
	q := []*TreeNode{root}
	for level := 1; len(q) > 0; level++ {
		temp := q
		q = nil
		sum := 0
		for _, node := range temp {
			sum += node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		if sum > maxSum {
			res, maxSum = level, sum
		}
	}

	return res
}
