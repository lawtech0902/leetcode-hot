/*
__author__ = 'robin-luo'
__date__ = '2025/06/16 11:13'
*/

package solution

func findAnagrams(s string, p string) []int {
	res := make([]int, 0)
	pCount := [26]int{}
	for _, ch := range p {
		pCount[ch-'a']++
	}

	l, r := 0, 0
	windowCharCount := [26]int{}
	for r < len(s) {
		windowCharCount[s[r]-'a']++
		if r-l+1 == len(p) {
			if isAnagram(windowCharCount, pCount) {
				res = append(res, l)
			}

			windowCharCount[s[l]-'a']--
			l++
		}

		r++
	}

	return res
}

func isAnagram(windowCharCount, pCount [26]int) bool {
	for i := 0; i < 26; i++ {
		if windowCharCount[i] != pCount[i] {
			return false
		}
	}

	return true
}
