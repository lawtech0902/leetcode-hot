/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 14:31'
*/

package solution

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}

	return res
}
