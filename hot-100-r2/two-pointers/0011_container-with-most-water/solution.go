/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 10:36'
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
