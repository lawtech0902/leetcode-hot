/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:02:38
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 17:09:08
 */

package solution

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
