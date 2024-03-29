/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 10:02'
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

func backtracking(target, curSum, startIndex int, candidates, track []int, res *[][]int) {
	if curSum == target {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
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
