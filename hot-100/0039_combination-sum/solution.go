/*
__author__ = 'robin-luo'
__date__ = '2023/03/13 19:15'
*/

package solution

func combinationSum(candidates []int, target int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(target, 0, 0, candidates, track, &res)
	return res
}

func backtracking(target, sum, startIndex int, candidates, track []int, res *[][]int) {
	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		track = append(track, candidates[i])
		sum += candidates[i]
		backtracking(target, sum, i, candidates, track, res)
		sum -= candidates[i]
		track = track[:len(track)-1]
	}
}
