/*
__author__ = 'robin-luo'
__date__ = '2025/06/13 11:08'
*/

package solution

func moveZeroes(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

		j++
	}
}
