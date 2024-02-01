/*
 * Author: robin-luo
 * Created time: 2024-02-01 16:08:34
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 16:12:16
 */

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	nums := inorder(root, []int{})
	minDiff := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		minDiff = min(minDiff, nums[i]-nums[i-1])
	}

	return minDiff
}

func inorder(root *TreeNode, nums []int) []int {
	if root == nil {
		return nums
	}

	nums = inorder(root.Left, nums)
	nums = append(nums, root.Val)
	nums = inorder(root.Right, nums)
	return nums
}
