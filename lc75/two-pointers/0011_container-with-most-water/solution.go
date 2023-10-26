/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-26 10:47:57
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 11:04:35
 * @Description:
 */

package solution

import "math"

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

		res = int(math.Max(float64(res), float64(area)))
	}

	return res
}
