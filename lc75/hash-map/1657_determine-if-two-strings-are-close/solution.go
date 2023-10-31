/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 10:19'
*/

package solution

import "sort"

func closeStrings(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	cnt1 := make([]int, 26)
	cnt2 := make([]int, 26)
	for _, ch := range word1 {
		cnt1[ch-'a']++
	}

	for _, ch := range word2 {
		cnt2[ch-'a']++
	}

	for i := 0; i < 26; i++ {
		if cnt1[i] == 0 && cnt2[i] != 0 || cnt1[i] != 0 && cnt2[i] == 0 {
			return false
		}
	}

	sort.Ints(cnt1)
	sort.Ints(cnt2)
	for i := 0; i < 26; i++ {
		if cnt1[i] != cnt2[i] {
			return false
		}
	}

	return true
}
