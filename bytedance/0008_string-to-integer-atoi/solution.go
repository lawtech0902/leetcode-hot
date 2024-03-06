/*
 * Author: robin-luo
 * Created time: 2024-03-05 11:24:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 13:45:55
 */

package solution

import "math"

func myAtoi(s string) int {
	sign := 1
	n := len(s)
	i := 0
	for i < n && s[i] == ' ' {
		i++
	}

	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	num := 0
	for i < n && '0' <= s[i] && s[i] <= '9' {
		num = 10*num + int(s[i]-'0')
		if sign*num < math.MinInt32 {
			return math.MinInt32
		} else if sign*num > math.MaxInt32 {
			return math.MaxInt32
		}

		i++
	}

	return sign * num
}
