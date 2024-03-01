/*
 * Author: robin-luo
 * Created time: 2024-02-29 16:50:23
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 16:56:29
 */

package solution

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	low, high := 1, x
	for low <= high {
		mid := (low + high) >> 1
		sqrt := x / mid
		if sqrt == mid {
			return mid
		} else if sqrt > mid {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}
