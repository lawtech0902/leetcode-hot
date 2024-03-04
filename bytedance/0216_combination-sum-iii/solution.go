/*
 * Author: robin-luo
 * Created time: 2024-03-01 11:28:40
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 13:49:04
 */

package solution

func combinationSum3(k int, n int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(k, n, 1, 0, track, &res)
	return res
}

func backtracking(k, n, startIndex, curSum int, track []int, res *[][]int) {
	if len(track) == k && curSum == n {
		temp := make([]int, k)
		copy(temp, track)
		*res = append(*res, temp)
	}

	if curSum > n {
		return
	}

	for i := startIndex; i <= 9; i++ {
		track = append(track, i)
		curSum += i
		backtracking(k, n, i+1, curSum, track, res)
		curSum -= i
		track = track[:len(track)-1]
	}
}
