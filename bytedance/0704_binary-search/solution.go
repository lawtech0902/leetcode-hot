/*
 * Author: robin-luo
 * Created time: 2024-03-05 16:51:42
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 16:52:54
 */

package solution

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
