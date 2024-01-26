/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 15:24'
*/

package solution

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return l
}
