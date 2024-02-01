/*
 * Author: robin-luo
 * Created time: 2024-01-31 14:29:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 14:32:00
 */

package solution

func minDistance(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)
	if l1 == 0 && l2 == 0 || word1 == word2 {
		return 0
	}

	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}

	for j := 0; j <= l2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
			}
		}
	}

	return dp[l1][l2]
}
