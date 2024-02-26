/*
 * Author: robin-luo
 * Created time: 2024-02-22 14:19:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 14:27:00
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

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

func dfs(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}

	if root.Val == targetSum {
		res++
	}

	res += dfs(root.Left, targetSum-root.Val)
	res += dfs(root.Right, targetSum-root.Val)
	return res
}
