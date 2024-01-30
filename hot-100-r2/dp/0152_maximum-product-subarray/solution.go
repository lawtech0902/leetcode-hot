/*
__author__ = 'robin-luo'
__date__ = '2024/01/30 10:42'
*/

package solution

import "math"

func maxProduct(nums []int) int {
	res := math.MinInt64
	iMax, iMin := 1, 1
	for _, num := range nums {
		if num < 0 {
			iMax, iMin = iMin, iMax
		}

		iMax = max(iMax*num, num)
		iMin = min(iMin*num, num)
		res = max(res, iMax)
	}

	return res
}
