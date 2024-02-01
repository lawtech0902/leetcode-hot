/*
 * Author: robin-luo
 * Created time: 2024-01-31 13:54:05
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 14:03:28
 */

package solution

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}

		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, l, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}
