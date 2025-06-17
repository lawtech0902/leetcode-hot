/*
__author__ = 'robin-luo'
__date__ = '2025/06/17 10:10'
*/

package solution

func minWindow(s string, t string) string {
	var res string
	if len(s) < len(t) {
		return res
	}

	hash := make([]int, 256)
	for i := range t {
		hash[t[i]]++
	}

	count := len(t)
	minLen := len(s) + 1
	l, r := 0, 0
	for r < len(s) {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}

		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if count == 0 && r-l+1 < minLen {
			minLen = r - l + 1
			res = s[l : r+1]
		}

		r++
	}

	return res
}
