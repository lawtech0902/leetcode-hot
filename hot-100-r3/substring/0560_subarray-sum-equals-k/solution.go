/*
__author__ = 'robin-luo'
__date__ = '2025/06/16 11:30'
*/

package solution

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

func subarraySum2(nums []int, k int) int {
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
