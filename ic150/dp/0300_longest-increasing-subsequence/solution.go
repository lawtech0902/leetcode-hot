/*
 * Author: robin-luo
 * Created time: 2024-01-30 21:38:59
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 09:59:01
 */

package solution

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	tail := 1
	d := make([]int, n+1)
	d[tail] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[tail] {
			tail++
			d[tail] = nums[i]
		} else {
			l, r := 1, tail
			pos := 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			d[pos+1] = nums[i]
		}
	}

	return tail
}
