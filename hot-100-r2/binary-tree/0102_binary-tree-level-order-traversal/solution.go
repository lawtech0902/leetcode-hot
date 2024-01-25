/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 19:52'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		tempArr := make([]int, len(queue))
		for i := range tempArr {
			node := queue[0]
			queue = queue[1:]
			tempArr[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}

func levelOrder1(root *TreeNode) [][]int {
	var (
		res [][]int
		dfs func(root *TreeNode, level int)
	)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if level >= len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)
	return res
}
