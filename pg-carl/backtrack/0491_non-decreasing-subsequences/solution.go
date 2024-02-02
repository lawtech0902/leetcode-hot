/*
 * Author: robin-luo
 * Created time: 2024-02-02 11:09:49
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 11:26:23
 */

package solution

func findSubsequences(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(nums, track, 0, &res)
	return res
}

func backtracking(nums, track []int, startIndex int, res *[][]int) {
	if len(track) > 1 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	history := [201]int{}
	for i := startIndex; i < len(nums); i++ {
		if len(track) > 0 && nums[i] < track[len(track)-1] || history[nums[i]+100] == 1 {
			continue
		}

		history[nums[i]+100] = 1
		track = append(track, nums[i])
		backtracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
