/*
__author__ = 'robin-luo'
__date__ = '2024/01/15 15:06'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var tempArr []int
		q2 := queue
		queue = nil
		for _, node := range queue {
			tempArr = append(tempArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				q2 = append(q2, node.Right)
			}
		}

		if level%2 == 1 {
			for i, n := 0, len(tempArr)-1; i < n/2; i++ {
				tempArr[i], tempArr[n-1-i] = tempArr[n-1-i], tempArr[i]
			}
		}

		res = append(res, tempArr)
	}

	return res
}
