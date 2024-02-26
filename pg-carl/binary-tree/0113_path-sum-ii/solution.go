/*
 * Author: robin-luo
 * Created time: 2024-02-22 14:15:57
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 14:19:13
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		res  [][]int
		path []int
		dfs  func(root *TreeNode, level, sum int)
	)

	dfs = func(root *TreeNode, level, sum int) {
		if root == nil {
			return
		}

		if level >= len(path) {
			path = append(path, root.Val)
		} else {
			path[level] = root.Val
		}

		sum -= root.Val

		if root.Left == nil && root.Right == nil && sum == 0 {
			tmp := make([]int, level+1)
			copy(tmp, path)
			res = append(res, tmp)
		}

		dfs(root.Left, level+1, sum)
		dfs(root.Right, level+1, sum)
	}

	dfs(root, 0, targetSum)
	return res
}
