/*
__author__ = 'robin-luo'
__date__ = '2023/12/11 10:28'
*/

package solution

import "strings"

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	strs := strings.Fields(s)
	var res strings.Builder
	for i := len(strs) - 1; i >= 0; i-- {
		res.WriteString(strs[i] + " ")
	}

	return strings.TrimSpace(res.String())
}

func reverseWords1(s string) string {
	slow, fast := 0, 0
	bytes := []byte(s)
	for len(bytes) > 0 && fast < len(bytes) && bytes[fast] == ' ' {
		fast++
	}

	for ; fast < len(bytes); fast++ {
		if fast-1 > 0 && bytes[fast-1] == bytes[fast] && bytes[fast] == ' ' {
			continue
		}

		bytes[slow] = bytes[fast]
		slow++
	}

	if slow-1 > 0 && bytes[slow-1] == ' ' {
		bytes = bytes[:slow-1]
	} else {
		bytes = bytes[:slow]
	}

	reverse(&bytes, 0, len(bytes)-1)
	i := 0
	for i < len(bytes) {
		j := i
		for j < len(bytes) && bytes[j] != ' ' {
			j++
		}

		reverse(&bytes, i, j-1)
		i = j
		i++
	}

	return string(bytes)
}

func reverse(bytes *[]byte, left, right int) {
	for left < right {
		(*bytes)[left], (*bytes)[right] = (*bytes)[right], (*bytes)[left]
		left++
		right--
	}
}
