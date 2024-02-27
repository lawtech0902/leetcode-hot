/*
 * Author: robin-luo
 * Created time: 2024-02-27 10:23:29
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 10:38:43
 */

package solution

import "strconv"

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	res := ""
	carryFlag := 0
	cur := 0
	for i >= 0 || j >= 0 {
		numA, numB := 0, 0
		if i >= 0 {
			numA = int(a[i] - '0')
		}

		if j >= 0 {
			numB = int(b[j] - '0')
		}

		cur = numA + numB + carryFlag
		carryFlag = 0
		if cur >= 2 {
			carryFlag = 1
			cur -= 2
		}

		curStr := strconv.Itoa(cur)
		res = curStr + res
		i++
		j--
	}

	if carryFlag == 1 {
		res = "1" + res
	}

	return res
}
