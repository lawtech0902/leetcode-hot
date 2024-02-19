/*
 * Author: robin-luo
 * Created time: 2024-02-19 14:57:09
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 15:01:43
 */

package solution

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += expandAroundCenter(s, i, i)
		res += expandAroundCenter(s, i, i+1)
	}

	return res
}

func expandAroundCenter(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}

	return count
}
