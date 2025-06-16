/*
__author__ = 'robin-luo'
__date__ = '2025/06/13 10:51'
*/

package solution

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	res := 0
	for num := range numSet {
		if !numSet[num-1] {
			curNum := num
			curStreak := 1
			for numSet[curNum+1] {
				curNum++
				curStreak++
			}

			res = max(res, curStreak)
		}
	}

	return res
}
