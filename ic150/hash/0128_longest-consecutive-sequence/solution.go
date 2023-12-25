/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 9:42'
*/

package solution

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	sort.Ints(nums)
	longestStreak, curStreak := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			if nums[i] == nums[i-1]+1 {
				curStreak++
			} else {
				longestStreak = max(longestStreak, curStreak)
				curStreak = 1
			}
		}
	}

	return max(longestStreak, curStreak)
}
