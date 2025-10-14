/*
__author__ = 'robin-luo'
__date__ = '2025/10/13 14:20'
*/

package lcr_001

import "math"

func divide(a, b int) int {
	if a == -1<<31 && b == -1 {
		return 1<<31 - 1
	}

	negative := (a < 0) != (b < 0)
	dvd, dvs := int(math.Abs(float64(a))), int(math.Abs(float64(b)))
	quotient := 0

	for dvd >= dvs {
		temp, multiple := dvs, 1
		for dvd >= temp<<1 {
			temp <<= 1
			multiple <<= 1
		}

		dvd -= temp
		quotient += multiple
	}

	if negative {
		quotient *= -1
	}

	if quotient < -1<<31 {
		return -1 << 31
	}

	if quotient > 1<<31-1 {
		return 1<<31 - 1
	}

	return quotient
}
