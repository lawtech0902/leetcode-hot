/*
 * Author: robin-luo
 * Created time: 2024-02-22 11:17:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 11:21:34
 */

package solution

import "strconv"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}

	dfs(root, strconv.Itoa(root.Val), &res)
	return res
}

func dfs(root *TreeNode, path string, res *[]string) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}

	if root.Left != nil {
		dfs(root.Left, path+"->"+strconv.Itoa(root.Left.Val), res)
	}

	if root.Right != nil {
		dfs(root.Right, path+"->"+strconv.Itoa(root.Right.Val), res)
	}
}
