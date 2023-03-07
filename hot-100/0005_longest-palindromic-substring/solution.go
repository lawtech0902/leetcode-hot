/*
__author__ = 'robin-luo'
__date__ = '2023/03/07 17:02'
*/

package solution

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	left, maxLen := 0, 1
	for i := 0; i < len(s); {
		if len(s)-i <= maxLen/2 {
			break
		}

		begin, end := i, i
		for end < len(s)-1 && s[end+1] == s[end] {
			end++
		}

		i = end + 1
		for end < len(s)-1 && begin > 0 && s[end+1] == s[begin-1] {
			end++
			begin--
		}

		if end+1-begin > maxLen {
			left = begin
			maxLen = end + 1 - begin
		}
	}

	return s[left : left+maxLen]
}
