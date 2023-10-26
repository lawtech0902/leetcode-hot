/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-26 11:07:44
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 11:11:33
 * @Description:
 */

package solution

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	count := 0
	for left < right {
		if nums[left]+nums[right] < k {
			left++
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			right--
			left++
			count++
		}
	}

	return count
}
