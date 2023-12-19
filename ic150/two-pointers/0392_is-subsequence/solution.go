/*
__author__ = 'robin-luo'
__date__ = '2023/12/19 9:32'
*/

package solution

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}
