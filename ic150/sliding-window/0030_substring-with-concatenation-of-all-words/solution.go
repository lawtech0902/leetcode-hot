/*
__author__ = 'robin-luo'
__date__ = '2023/12/20 9:41'
*/

package solution

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 {
		return nil
	}

	wordLen := len(words[0])
	totalCount := len(words)
	totalLen := wordLen * totalCount
	if totalLen > len(s) {
		return nil
	}

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	res := make([]int, 0)
	for i := 0; i <= len(s)-totalLen; i++ {
		seen := make(map[string]int)
		j := 0
		for j < totalCount {
			word := s[i+j*wordLen : i+j*wordLen+wordLen]
			if count, ok := wordMap[word]; !ok || seen[word] >= count {
				break
			}

			seen[word]++
			j++
		}

		if j == totalCount {
			res = append(res, i)
		}
	}

	return res
}
