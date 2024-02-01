/*
 * Author: robin-luo
 * Created time: 2024-02-01 19:39:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 19:48:30
 */

package solution

func combinationSum3(k, n int) [][]int {
	var (
		res   [][]int
		track []int
	)

	backtracking(n, k, 1, track, &res)
	return res
}

func backtracking(n, k, startIndex int, track []int, res *[][]int) {
	if len(track) == k {
		sum := 0
		temp := make([]int, k)
		copy(temp, track)
		for _, v := range track {
			sum += v
		}

		if sum == n {
			*res = append(*res, temp)
		}

		return
	}

	for i := startIndex; i <= 9-(k-len(track))+1; i++ {
		track = append(track, i)
		backtracking(n, k, i+1, track, res)
		track = track[:len(track)-1]
	}
}
