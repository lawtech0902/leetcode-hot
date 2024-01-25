/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 10:32'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// basic version
func pathSum(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}

	res = dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(node *TreeNode, targetSum int) int {
	res := 0
	if node == nil {
		return res
	}

	if node.Val == targetSum {
		res++
	}

	res += dfs(node.Left, targetSum-node.Val)
	res += dfs(node.Right, targetSum-node.Val)
	return res
}

// advanced version
func pathSum1(root *TreeNode, targetSum int) int {
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
