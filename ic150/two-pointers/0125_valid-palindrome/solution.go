/*
__author__ = 'robin-luo'
__date__ = '2023/12/15 14:36'
*/

package solution

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		if !('a' <= s[left] && s[left] <= 'z' || '0' <= s[left] && s[left] <= '9') {
			left++
			continue
		}

		if !('a' <= s[right] && s[right] <= 'z' || '0' <= s[right] && s[right] <= '9') {
			right--
			continue
		}

		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
