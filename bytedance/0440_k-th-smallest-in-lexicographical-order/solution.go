/*
 * Author: robin-luo
 * Created time: 2024-03-04 14:47:38
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 17:07:33
 */

package solution

// https://blog.csdn.net/qq_36389060/article/details/124107294
func findKthNumber(n, k int) int {
	p := 1
	prefix := 1
	for p < k {
		count := getCount(prefix, n)
		if p+count > k {
			prefix *= 10
			p++
		} else {
			prefix++
			p += count
		}
	}

	return prefix
}

func getCount(prefix, n int) int {
	cur := prefix
	next := prefix + 1
	count := 0
	for cur <= n {
		count += min(n+1, next) - cur
		cur *= 10
		next *= 10
	}

	return count
}
