/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 17:36'
*/

package solution

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}

	sArr := strings.Split(s, " ")
	return len(sArr[len(sArr)-1])
}
