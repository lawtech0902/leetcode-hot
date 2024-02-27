/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:41:31
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 17:43:35
 */

package solution

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i; j >= 0; j-- {
			if dp[j] && inWordDict(s[j:i+1], wordDict) {
				dp[i+1] = true
			}
		}
	}

	return dp[n]
}

func inWordDict(s string, wordDict []string) bool {
	for _, word := range wordDict {
		if s == word {
			return true
		}
	}

	return false
}
