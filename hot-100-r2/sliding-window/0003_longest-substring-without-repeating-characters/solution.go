/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 15:45'
*/

package solution

func lengthOfLongestSubstring(s string) int {
	location := [128]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxLen
}
