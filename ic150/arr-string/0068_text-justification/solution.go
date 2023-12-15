/*
__author__ = 'robin-luo'
__date__ = '2023/12/15 10:33'
*/

package solution

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var res []string
	right, n := 0, len(words)
	for {
		left := right
		sumLen := 0
		for right < n && sumLen+len(words[right])+right-left <= maxWidth {
			sumLen += len(words[right])
			right++
		}

		if right == n {
			s := strings.Join(words[left:], " ")
			res = append(res, s+blank(maxWidth-len(s)))
			return res
		}

		numWords := right - left
		numSpaces := maxWidth - sumLen

		if numWords == 1 {
			res = append(res, words[left]+blank(numSpaces))
			continue
		}

		avgSpaces := numSpaces / (numWords - 1)
		extraSpaces := numSpaces % (numWords - 1)
		s1 := strings.Join(words[left:left+extraSpaces+1], blank(avgSpaces+1))
		s2 := strings.Join(words[left+extraSpaces+1:right], blank(avgSpaces))
		res = append(res, s1+blank(avgSpaces)+s2)
	}
}

func blank(n int) string {
	return strings.Repeat(" ", n)
}
