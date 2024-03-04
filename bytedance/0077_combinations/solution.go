/*
 * Author: robin-luo
 * Created time: 2024-03-01 11:11:49
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 11:14:51
 */

package solution

func combine(n, k int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(n, k, 1, track, &res)
	return res
}

func backtracking(n, k, startIndex int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		*res = append(*res, temp)
	}

	if len(track)+n-startIndex+1 < k {
		return
	}

	for i := startIndex; i <= n; i++ {
		track = append(track, i)
		backtracking(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
