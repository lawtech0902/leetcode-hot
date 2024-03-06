/*
 * Author: robin-luo
 * Created time: 2024-03-05 15:12:45
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 15:56:54
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 中序有序
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}

		root = root.Right
	}
}
