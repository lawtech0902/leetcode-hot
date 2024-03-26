/*
 * Author: robin-luo
 * Created time: 2024-03-19 15:30:08
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 15:41:58
 */

package solution

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	for _, num := range nums {
		if num%2 == 0 {
			res[l] = num
			l++
		} else {
			res[r] = num
			r--
		}
	}

	return res
}
