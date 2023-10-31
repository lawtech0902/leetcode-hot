/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-27 10:37:03
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-27 13:41:47
 * @Description:
 */

package solution

import "math"

func longestOnes(nums []int, k int) int {
	left, lsum, rsum, res := 0, 0, 0, 0
	for right, num := range nums {
		rsum += 1 - num
		for lsum < rsum-k {
			lsum += 1 - nums[left]
			left++
		}

		res = int(math.Max(float64(res), float64(right-left+1)))
	}

	return res
}
