/*
 * Author: robin-luo
 * Created time: 2024-03-19 16:34:40
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 16:35:36
 */

package solution

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
