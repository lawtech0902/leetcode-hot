/*
 * Author: robin-luo
 * Created time: 2024-02-29 13:51:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 13:56:24
 */

package solution

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	// 从后往前找第一个相邻升序对A[i] < A[j]
	i, j := n-2, n-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 从后往前找[j:]中第一个大于A[i]的A[k], 然后交换
	k := n - 1
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}

		nums[i], nums[k] = nums[k], nums[i]
	}

	// 逆序A[j:]
	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
