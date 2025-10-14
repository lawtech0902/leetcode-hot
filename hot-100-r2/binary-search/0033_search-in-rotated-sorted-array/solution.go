/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 16:13'
*/

package solution

func search(nums []int, target int) int {
	size := len(nums)
	l, r := 0, size-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		// 左半部分有序
		if nums[mid] >= nums[l] {
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
