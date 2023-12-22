/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 11:34'
*/

package solution

import "strings"

func wordPattern(pattern, s string) bool {
	sArr := strings.Split(s, " ")
	if len(pattern) != len(sArr) {
		return false
	}

	patternMap := make(map[byte]string)
	mappedWord := make(map[string]bool)
	for i := 0; i < len(pattern); i++ {
		matchedWord, ok := patternMap[pattern[i]]
		if !ok {
			if mapped := mappedWord[sArr[i]]; mapped {
				return false
			}

			patternMap[pattern[i]] = sArr[i]
			mappedWord[sArr[i]] = true
			continue
		}

		if sArr[i] != matchedWord {
			return false
		}
	}

	return true
}
