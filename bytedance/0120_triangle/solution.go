/*
 * Author: robin-luo
 * Created time: 2024-02-28 14:16:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 15:00:12
 */

package solution

// 自底向上+空间优化
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}
