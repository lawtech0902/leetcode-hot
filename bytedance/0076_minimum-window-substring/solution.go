/*
 * Author: robin-luo
 * Created time: 2024-02-27 15:39:51
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 15:52:12
 */

package solution

func minWindow(s, t string) string {
	if len(s) < len(t) {
		return ""
	}

	hash := make([]int, 256)
	for i := range t {
		hash[t[i]]++
	}

	l, count, minLen, res := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}

		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}

		if count == 0 && minLen > r-l+1 {
			minLen = r - l + 1
			res = s[l : r+1]
		}
	}

	return res
}
