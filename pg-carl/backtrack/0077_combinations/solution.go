/*
 * Author: robin-luo
 * Created time: 2024-02-01 19:16:31
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 19:34:52
 */

package solution

func combine(n, k int) [][]int {
	var (
		res   [][]int
		track []int
	)

	backtracking(n, k, 1, track, &res)
	return res
}

func backtracking(n, k, startIndex int, track []int, res *[][]int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		*res = append(*res, track)
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
