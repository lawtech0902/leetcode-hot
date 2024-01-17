/*
__author__ = 'robin-luo'
__date__ = '2024/01/17 11:29'
*/

package solution

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}

	l, count, maxLen, res := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}

		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if count == 0 && maxLen > r-l+1 {
			maxLen = r - l + 1
			res = s[l : r+1]
		}
	}

	return res
}
