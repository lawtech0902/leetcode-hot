/*
 * Author: robin-luo
 * Created time: 2024-03-01 11:18:01
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 11:28:01
 */

package solution

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)

	slices.Sort(candidates)
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
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}

		track = append(track, candidates[i])
		curSum += candidates[i]
		backtracking(candidates, track, i+1, curSum, target, res)
		curSum -= candidates[i]
		track = track[:len(track)-1]
	}
}
