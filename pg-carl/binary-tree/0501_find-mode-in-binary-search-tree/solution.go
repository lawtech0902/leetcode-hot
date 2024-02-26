/*
 * Author: robin-luo
 * Created time: 2024-02-22 15:52:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:59:35
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var (
		stack  []*TreeNode
		memo   = make(map[int]int)
		res    []int
		maxVal int
		cur    = root
	)

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			memo[cur.Val]++
			if maxVal < memo[cur.Val] {
				maxVal = memo[cur.Val]
				res = nil
				res = append(res, cur.Val)
			} else if maxVal == memo[cur.Val] {
				res = append(res, cur.Val)
			}

			cur = cur.Right
		}
	}

	return res
}
