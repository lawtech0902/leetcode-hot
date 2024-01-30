/*
__author__ = 'robin-luo'
__date__ = '2024/01/30 16:30'
*/

package solution

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
