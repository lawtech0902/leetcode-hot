/*
 * Author: robin-luo
 * Created time: 2024-03-01 11:00:53
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 11:04:01
 */

package solution

func combinationSum(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(candidates, track, 0, 0, target, &res)
	return res
}

func backtracking(candidates, track []int, startIndex, curSum, target int, res *[][]int) {
	if curSum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	if curSum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		track = append(track, candidates[i])
		curSum += candidates[i]
		backtracking(candidates, track, i, curSum, target, res)
		curSum -= candidates[i]
		track = track[:len(track)-1]
	}
}
