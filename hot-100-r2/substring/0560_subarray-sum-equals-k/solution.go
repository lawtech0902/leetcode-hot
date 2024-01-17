/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 16:53'
*/

package solution

func subarraySum(nums []int, k int) int {
	count := 0
	preSum := map[int]int{0: 1}
	sum := 0
	for _, num := range nums {
		sum += num
		if val, ok := preSum[sum-k]; ok {
			count += val
		}

		preSum[sum]++
	}

	return count
}
