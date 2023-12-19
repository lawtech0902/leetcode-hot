/*
__author__ = 'robin-luo'
__date__ = '2023/12/19 13:40'
*/

package solution

func lengthOfLongestSubstring(s string) int {
	location := [128]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i := 0; i < len(s); i++ {
		if location[s[i]] > left {
			left = location[s[i]] + 1
		} else {
			maxLen = i + 1 - left
		}

		location[s[i]] = i
	}

	return maxLen
}
