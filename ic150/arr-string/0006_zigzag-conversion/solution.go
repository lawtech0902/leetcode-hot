/*
__author__ = 'robin-luo'
__date__ = '2023/12/11 11:18'
*/

package solution

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	res := make([]strings.Builder, numRows)
	i, flag := 0, -1
	for _, c := range s {
		res[i].WriteRune(c)
		if i == 0 || i == numRows-1 {
			flag *= -1
		}

		i += flag
	}

	var resBuilder strings.Builder
	for _, sb := range res {
		resBuilder.WriteString(sb.String())
	}

	return resBuilder.String()
}
