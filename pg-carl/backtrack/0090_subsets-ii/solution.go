/*
 * Author: robin-luo
 * Created time: 2024-02-02 10:58:38
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 11:09:20
 */

package solution

import "slices"

func subsetsWithDup(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	slices.Sort(nums)
	backtracking(nums, track, 0, &res)
	return res
}

func backtracking(nums, track []int, startIndex int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		}

		track = append(track, nums[i])
		backtracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
