/*
__author__ = 'robin-luo'
__date__ = '2024/01/17 16:21'
*/

package solution

func maxSubArray(nums []int) int {
	res := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		res = max(res, sum)
	}

	return res
}
