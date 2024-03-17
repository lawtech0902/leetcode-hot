/*
 * Author: robin-luo
 * Created time: 2024-03-12 14:22:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 16:50:21
 */

package solution

func search(nums []int, target int) int {
	// 获取数组长度
	n := len(nums)
	// 初始左右边界
	l, r := 0, n-1
	// 开始二分查找
	for l <= r {
		// 若左边界对应的值即为目标值，直接返回左边界下标
		if nums[l] == target {
			return l
		}

		// 计算中间值的下标
		mid := (l + r) >> 1
		// 若中间值为目标值，将右边界移动到中间值，继续查找
		if nums[mid] == target {
			r = mid
		} else if nums[0] < nums[mid] { // 左半部分有序
			// 若目标值在左半部分有序区间内，调整右边界
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				// 否则调整左边界
				l = mid + 1
			}
		} else if nums[0] > nums[mid] { // 右半部分有序
			// 若目标值在右半部分有序区间内，调整左边界
			if nums[mid] < target && target < nums[n-1] {
				l = mid + 1
			} else {
				// 否则调整右边界
				r = mid - 1
			}
		} else {
			// 中间值和最左值相等，递增左边界以跳过重复值
			l++
		}
	}

	// 未找到目标值，返回 -1
	return -1
}
