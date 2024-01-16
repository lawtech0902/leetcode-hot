/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 10:08'
*/

package solution

import (
	"slices"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	slices.Sort(nums)
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

func longestConsecutive1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			curNum := num
			curStreak := 1
			for numSet[curNum+1] {
				curNum++
				curStreak++
			}

			longestStreak = max(longestStreak, curStreak)
		}
	}

	return longestStreak
}
