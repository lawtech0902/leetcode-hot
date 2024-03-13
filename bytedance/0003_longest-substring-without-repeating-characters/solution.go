/*
 * Author: robin-luo
 * Created time: 2024-02-27 16:14:01
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-12 17:41:55
 */

package solution

func lengthOfLongestSubstring(s string) int {
	location := [128]int{}
	for i := range location {
		location[i] = -1
	}

	maxLen, left := 0, 0
	for i, ch := range s {
		if location[ch] >= left {
			left = location[ch] + 1
		} else if i+1-left > maxLen {
			maxLen = i + 1 - left
		}

		location[ch] = i
	}

	return maxLen
}
