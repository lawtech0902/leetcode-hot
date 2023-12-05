/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 11:21'
*/

package solution

func removeElements(nums []int, val int) int {
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		} else {
			nums[cur] = nums[i]
			cur++
		}
	}

	return cur
}
