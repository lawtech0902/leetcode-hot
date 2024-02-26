/*
 * Author: robin-luo
 * Created time: 2024-02-22 14:39:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 15:10:09
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	index := getMaxValIndex(nums)
	root := &TreeNode{Val: nums[index]}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func getMaxValIndex(nums []int) int {
	index := 0
	for i, v := range nums {
		if v > nums[index] {
			index = i
		}
	}

	return index
}
