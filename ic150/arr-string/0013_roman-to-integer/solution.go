/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 16:37'
*/

package solution

func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]
		sign := 1
		if temp < last {
			sign = -1
		}

		res += sign * temp
		last = temp
	}

	return res
}
