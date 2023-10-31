/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-27 14:02:07
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-30 10:49:38
 * @Description:
 */

package solution

func longestSubarray(nums []int) int {
	left, index, k := 0, 0, 1
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			k--
		} else {
			res = max(res, i-1-left)
			left = index + 1
		}

		index = i
	}

	res = max(res, len(nums)-1-left)
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
