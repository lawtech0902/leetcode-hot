/*
 * Author: robin-luo
 * Created time: 2024-02-28 09:37:28
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 09:38:44
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
