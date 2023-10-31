/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-27 09:20:12
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-27 09:45:03
 * @Description:
 */

package solution

import "math"

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	res := float64(sum)
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		res = math.Max(res, float64(sum))
	}

	return res / float64(k)
}
