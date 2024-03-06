/*
 * Author: robin-luo
 * Created time: 2024-03-06 10:09:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-06 10:15:09
 */

package solution

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res, area := 0, 0
	for l < r {
		if height[l] < height[r] {
			area = height[l] * (r - l)
			l++
		} else {
			area = height[r] * (r - l)
			r--
		}

		res = max(res, area)
	}

	return res
}
