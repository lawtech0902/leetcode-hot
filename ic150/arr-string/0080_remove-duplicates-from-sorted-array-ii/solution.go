/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 15:14'
*/

package solution

func removeDuplicates(nums []int) int {
	len := 0
	for _, num := range nums {
		if len < 2 || num != nums[len-2] {
			nums[len] = num
			len++
		}
	}

	return len
}
