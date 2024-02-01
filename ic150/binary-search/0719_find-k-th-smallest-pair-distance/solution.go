/*
 * Author: robin-luo
 * Created time: 2024-02-01 14:03:10
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 14:23:16
 */

package solution

import (
	"slices"
)

/*
可以对数对的绝对差值(即数对距离)做二分，首先排序数组，定义l和r：
开始二分，计算差值绝对值的中间值mid；
使用双指针计算有多少对元素cnt小于该中间值；
如果cnt比k要大，说明k在二分范围的左边，重置r；否则重置l；
*/

func smallestDistancePair(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	left := 0
	// 最大差值
	right := nums[n-1] - nums[0]
	for left <= right {
		mid := (left + right) >> 1
		count := 0
		for i, j := 0, 0; i < n; i++ {
			for j < n && nums[j]-nums[i] <= mid {
				j++
			}

			count += j - i - 1
		}

		if count < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
