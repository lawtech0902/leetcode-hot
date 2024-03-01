/*
 * Author: robin-luo
 * Created time: 2024-02-29 16:15:01
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 16:16:58
 */

package solution

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	maxSize := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}

			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}

			maxSize = max(maxSize, dp[i][j])
		}
	}

	return maxSize * maxSize
}
