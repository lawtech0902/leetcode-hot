/*
 * Author: robin-luo
 * Created time: 2024-02-27 16:14:01
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 16:16:00
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
