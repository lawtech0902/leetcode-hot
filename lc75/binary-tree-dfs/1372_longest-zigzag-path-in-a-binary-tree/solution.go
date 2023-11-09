/*
__author__ = 'robin-luo'
__date__ = '2023/11/06 10:02'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	res := 0
	dfs(root, false, 0, &res)
	dfs(root, true, 0, &res)
	return res
}

func dfs(node *TreeNode, dir bool, len int, res *int) {
	*res = max(*res, len)
	if !dir {
		if node.Left != nil {
			dfs(node.Left, true, len+1, res)
		}

		if node.Right != nil {
			dfs(node.Right, false, 1, res)
		}
	} else {
		if node.Right != nil {
			dfs(node.Right, false, len+1, res)
		}

		if node.Left != nil {
			dfs(node.Left, true, 1, res)
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
