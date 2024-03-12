/*
 * Author: robin-luo
 * Created time: 2024-03-06 10:49:03
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-06 10:59:53
 */

package solution

// 暴力
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}

	return count
}

func subarraySum1(nums []int, k int) int {
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
