/*
__author__ = 'robin-luo'
__date__ = '2025/10/14 16:24'
*/

package lcr_002

import "strconv"

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	res := ""
	carryFlag := 0
	cur := 0
	for i >= 0 || j >= 0 {
		intA, intB := 0, 0
		if i >= 0 {
			intA = int(a[i] - '0')
		}

		if j >= 0 {
			intB = int(b[j] - '0')
		}

		cur = intA + intB + carryFlag
		carryFlag = 0
		if cur >= 2 {
			carryFlag = 1
			cur -= 2
		}

		curStr := strconv.Itoa(cur)
		res = curStr + res
		i--
		j--
	}

	return res
}
