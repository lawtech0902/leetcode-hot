/*
__author__ = 'robin-luo'
__date__ = '2024/01/30 16:44'
*/

package solution

// 三指针 交换
func sortColors(nums []int) {
	p0, p1 := 0, 0
	p2 := len(nums) - 1
	for p1 <= p2 {
		if nums[p1] == 0 {
			nums[p0], nums[p1] = nums[p1], nums[p0]
			p0++
			p1++
		} else if nums[p1] == 2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
			p2--
		} else {
			p1++
		}
	}
}

// 双指针 赋值
func sortColors1(nums []int) {
	p0, p1 := 0, 0
	for i := range nums {
		temp := nums[i]
		nums[i] = 2
		if temp < 2 {
			nums[p1] = 1
			p1++
		}

		if temp < 1 {
			nums[p0] = 0
			p0++
		}
	}
}
