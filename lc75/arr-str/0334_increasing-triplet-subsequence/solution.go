/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-25 11:39:58
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 09:29:02
 * @Description:
 */

package solution

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	low, medium := nums[0], math.MaxInt32
	for _, v := range nums {
		if v > medium {
			return true
		}

		if v < medium && v > low {
			medium = v
		}

		if v < low {
			low = v
		}
	}

	return false
}
