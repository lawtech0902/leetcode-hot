/*
 * Author: robin-luo
 * Created time: 2024-02-02 09:37:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 09:58:03
 */

package solution

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)

	slices.Sort(candidates)
	backtracking(target, 0, 0, candidates, track, &res)
	return res
}

func backtracking(target, curSum, startIndex int, candidates, track []int, res *[][]int) {
	if curSum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	if curSum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		curSum += candidates[i]
		backtracking(target, curSum, i+1, candidates, track, res)
		curSum -= candidates[i]
		track = track[:len(track)-1]
	}
}
