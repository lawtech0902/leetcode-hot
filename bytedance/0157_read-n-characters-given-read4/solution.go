/*
 * Author: robin-luo
 * Created time: 2024-03-17 19:52:19
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-17 19:55:53
 */

package solution

// fake read4 func
func read4(buf []byte) int {
	return 0
}

func read(buf []byte, n int) int {
	total := 0
	temp := make([]byte, 4)
	for total < n {
		count := read4(temp)
		if count == 0 {
			break
		}

		for i := 0; i < count && total < n; i++ {
			buf[total] = temp[i]
			total++
		}
	}

	return total
}
