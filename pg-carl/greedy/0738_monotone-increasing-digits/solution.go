/*
 * Author: robin-luo
 * Created time: 2024-02-04 10:33:25
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 10:53:53
 */

package solution

import "strconv"

func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	size := len(ss)
	if size <= 1 {
		return n
	}

	for i := size - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < n; i++ {
				ss[j] = '9'
			}
		}
	}

	res, _ := strconv.Atoi(string(ss))
	return res
}
