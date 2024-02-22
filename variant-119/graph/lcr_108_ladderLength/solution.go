/*
 * Author: robin-luo
 * Created time: 2024-02-20 14:55:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 15:01:48
 */

package solution

func ladderLength(beginWord, endWord string, wordList []string) int {
	wordMap := getWordMap(wordList, beginWord)
	q := []string{beginWord}
	depth := 0
	for len(q) > 0 {
		depth++
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			word := q[0]
			q = q[1:]
			candidates := getCandidates(word)
			for _, candidate := range candidates {
				if _, exist := wordMap[candidate]; exist {
					if candidate == endWord {
						return depth + 1
					}

					delete(wordMap, candidate)
					q = append(q, candidate)
				}
			}
		}
	}

	return 0
}

func getWordMap(wordList []string, beginWord string) map[string]int {
	wordMap := make(map[string]int)
	for i, word := range wordList {
		if _, exist := wordMap[word]; !exist {
			if word != beginWord {
				wordMap[word] = i
			}
		}
	}

	return wordMap
}

func getCandidates(word string) []string {
	var res []string
	for i := 0; i < 26; i++ {
		for j := 0; j < len(word); j++ {
			if word[j] != byte(int('a')+i) {
				res = append(res, word[:j]+string(int('a')+i)+word[j+1:])
			}
		}
	}

	return res
}
