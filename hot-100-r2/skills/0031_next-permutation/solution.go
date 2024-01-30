/*
__author__ = 'robin-luo'
__date__ = '2024/01/30 17:08'
*/

package solution

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// find nums[i] < nums[j]
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	// 非最后一个排列
	if i >= 0 {
		// find nums[i] < nums[k]
		for nums[i] >= nums[k] {
			k--
		}

		// swap nums[i], nums[k]
		nums[i], nums[k] = nums[k], nums[i]
	}

	// reverse nums[j:end]
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
