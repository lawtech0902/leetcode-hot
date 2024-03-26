/*
 * Author: robin-luo
 * Created time: 2024-03-20 13:53:57
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-20 13:56:41
 */

package solution

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x *= -1
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = temp + res*10
		x = x / 10
	}

	res = sign * res
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
