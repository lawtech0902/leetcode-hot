/*
 * Author: robin-luo
 * Created time: 2024-02-07 16:54:14
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-07 17:13:16
 */

package solution

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	if int(math.Abs(float64(target))) > sum || (sum+target)%2 == 1 {
		return 0
	}

	bagSize := (sum + target) / 2
	dp := make([]int, bagSize+1)
	dp[0] = 1
	for _, num := range nums {
		for j := bagSize; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[bagSize]
}
