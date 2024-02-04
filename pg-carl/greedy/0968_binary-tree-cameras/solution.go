/*
 * Author: robin-luo
 * Created time: 2024-02-04 10:58:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 19:47:02
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	var (
		res int
		dfs func(node *TreeNode) int
	)

	dfs = func(node *TreeNode) int {
		// 空节点，该节点有覆盖
		if node == nil {
			return 2
		}

		// 后序遍历
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left == 2 && right == 2 {
			// 情况1
			// 左右节点都有覆盖
			return 0
		} else if left == 0 || right == 0 {
			// 情况2
			// left == 0 && right == 0 左右节点无覆盖
			// left == 1 && right == 0 左节点有摄像头，右节点无覆盖
			// left == 0 && right == 1 左节点有无覆盖，右节点摄像头
			// left == 0 && right == 2 左节点无覆盖，右节点覆盖
			// left == 2 && right == 0 左节点覆盖，右节点无覆盖
			res++
			return 1
		} else {
			// 情况3
			// left == 1 && right == 2 左节点有摄像头，右节点有覆盖
			// left == 2 && right == 1 左节点有覆盖，右节点有摄像头
			// left == 1 && right == 1 左右节点都有摄像头
			// 其他情况前段代码均已覆盖
			return 2
		}
	}

	if dfs(root) == 0 {
		res++
	}

	return res
}
