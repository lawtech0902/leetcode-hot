/*
__author__ = 'robin-luo'
__date__ = '2023/12/19 10:57'
*/

package solution

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, area := 0, 0
	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		res = max(res, area)
	}

	return res
}
