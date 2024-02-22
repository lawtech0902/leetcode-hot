/*
 * Author: robin-luo
 * Created time: 2024-02-21 19:57:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 20:07:02
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, pathSum int) int
	dfs = func(root *TreeNode, pathSum int) int {
		if root == nil {
			return 0
		}

		pathSum = pathSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return pathSum
		}

		return dfs(root.Left, pathSum) + dfs(root.Right, pathSum)
	}

	return dfs(root, 0)
}
