/*
 * Author: robin-luo
 * Created time: 2024-03-17 20:46:24
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-17 20:49:28
 */

package solution

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l]
}
