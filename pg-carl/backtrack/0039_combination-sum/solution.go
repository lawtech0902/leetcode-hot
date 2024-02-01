/*
 * Author: robin-luo
 * Created time: 2024-02-01 20:12:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 20:16:02
 */

package solution

func combinationSum(candidates []int, target int) [][]int {
	var (
		res   [][]int
		track []int
	)

	backtracking(target, 0, 0, candidates, track, &res)
	return res
}

func backtracking(target, curSum, startIndex int, candidates, track []int, res *[][]int) {
	if curSum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	} else if curSum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		track = append(track, candidates[i])
		curSum += candidates[i]
		backtracking(target, curSum, i, candidates, track, res)
		curSum -= candidates[i]
		track = track[:len(track)-1]
	}
}
