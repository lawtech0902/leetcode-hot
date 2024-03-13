/*
__author__ = 'robin-luo'
__date__ = '2023/03/10 15:50'
*/

package solution

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= target {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if nums[mid] <= nums[r] {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
