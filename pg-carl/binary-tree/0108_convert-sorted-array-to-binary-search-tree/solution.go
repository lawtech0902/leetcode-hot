/*
 * Author: robin-luo
 * Created time: 2024-02-22 16:54:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 16:56:09
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	mid := (l + r) >> 1
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, l, mid-1)
	root.Right = helper(nums, mid+1, r)
	return root
}
