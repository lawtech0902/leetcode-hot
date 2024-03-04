/*
 * Author: robin-luo
 * Created time: 2024-03-04 10:01:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 10:26:40
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(root, targetSum, track, &res)
	return res
}

func backtracking(root *TreeNode, targetSum int, track []int, res *[][]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			track = append(track, root.Val)
			temp := make([]int, len(track))
			copy(temp, track)
			*res = append(*res, temp)
		}

		return
	}

	track = append(track, root.Val)
	backtracking(root.Left, targetSum-root.Val, track, res)
	backtracking(root.Right, targetSum-root.Val, track, res)
	track = track[:len(track)-1]
}
