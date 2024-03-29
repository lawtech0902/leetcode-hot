/*
 * Author: robin-luo
 * Created time: 2024-02-21 20:07:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 17:12:34
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var res int
	if root == nil {
		return 0
	}

	res += dfs(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

func dfs(root *TreeNode, targetSum int) int {
	var res int
	if root == nil {
		return 0
	}

	if root.Val == targetSum {
		res++
	}

	res += dfs(root.Left, targetSum-root.Val)
	res += dfs(root.Right, targetSum-root.Val)
	return res
}
