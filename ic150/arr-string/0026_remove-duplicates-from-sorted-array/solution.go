/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 11:36'
*/

package solution

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	cur := 0
	for i := 1; i < size; i++ {
		if nums[i] != nums[i-1] {
			cur++
			nums[cur] = nums[i]
		}
	}

	return cur + 1
}
