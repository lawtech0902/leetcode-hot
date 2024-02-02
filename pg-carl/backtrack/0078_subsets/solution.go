/*
 * Author: robin-luo
 * Created time: 2024-02-02 10:41:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 10:54:50
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
