/*
 * Author: robin-luo
 * Created time: 2024-02-07 17:13:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-07 17:24:58
 */

package solution

func findMaxForm(strs []string, m, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		oneNum, zeroNum := 0, 0
		for _, ch := range str {
			if ch == '1' {
				oneNum++
			} else {
				zeroNum++
			}
		}

		for i := m; i >= zeroNum; i-- {
			for j := n; j >= oneNum; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}

	return dp[m][n]
}
