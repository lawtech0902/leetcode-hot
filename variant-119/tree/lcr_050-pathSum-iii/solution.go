/*
 * Author: robin-luo
 * Created time: 2024-02-21 20:07:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 20:13:40
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	var dfs func(root *TreeNode, curSum int)
	dfs = func(root *TreeNode, curSum int) {
		if root == nil {
			return
		}

		curSum += root.Val
		res += prefixSum[curSum-targetSum]
		prefixSum[curSum]++
		dfs(root.Left, curSum)
		dfs(root.Right, curSum)
		prefixSum[curSum]--
	}

	dfs(root, 0)
	return res
}
