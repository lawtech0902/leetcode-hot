/*
 * Author: robin-luo
 * Created time: 2024-03-05 10:53:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 11:05:17
 */

package solution

func longestConsecutive(nums []int) int {
	// 创建一个map，用于存放nums中的元素，key为元素值，value为true
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	// 初始化结果为0
	res := 0
	// 遍历map中的元素
	for num := range numSet {
		// 如果当前元素的前一个元素不存在map中，则开始计算当前元素的连续序列
		if !numSet[num-1] {
			curNum := num
			curStreak := 1
			// 继续判断当前元素的后一个元素是否存在于map中，若存在，则继续累加
			for numSet[curNum+1] {
				curNum++
				curStreak++
			}

			// 更新结果
			res = max(res, curStreak)
		}
	}

	return res
}
