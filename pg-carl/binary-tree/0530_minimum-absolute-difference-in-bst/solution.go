/*
 * Author: robin-luo
 * Created time: 2024-02-22 15:26:51
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:46:19
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var prev *TreeNode
	minDiff := math.MaxInt64
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		if prev != nil && node.Val-prev.Val < minDiff {
			minDiff = node.Val - prev.Val
		}

		prev = node
		traversal(node.Right)
	}
	traversal(root)
	return minDiff
}

func getMinimumDifference1(root *TreeNode) int {
	stack := []*TreeNode{}
	cur := root
	var pre *TreeNode
	minDiff := math.MaxInt64
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pre != nil {
				minDiff = min(minDiff, cur.Val-pre.Val)
			}

			pre = cur
			cur = cur.Right
		}
	}

	return minDiff
}
