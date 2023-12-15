/*
__author__ = 'robin-luo'
__date__ = '2023/12/15 10:26'
*/

package solution

func strStr(haystack string, needle string) int {
	l1, l2 := len(haystack), len(needle)
	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}

	return -1
}
