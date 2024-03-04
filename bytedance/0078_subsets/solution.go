/*
 * Author: robin-luo
 * Created time: 2024-03-01 16:14:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 16:16:18
 */

package solution

func subsets(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(nums, track, 0, &res)
	return res
}

func backtracking(nums, track []int, startIndex int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)

	for i := startIndex; i < len(nums); i++ {
		track = append(track, nums[i])
		backtracking(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}
