/*
 * Author: robin-luo
 * Created time: 2024-02-27 16:17:25
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 16:32:03
 */

package solution

func permute(nums []int) [][]int {
	var (
		track []int
		res   [][]int
	)

	backtracking(nums, track, &res)
	return res
}

func backtracking(nums, track []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		track = append(track, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backtracking(nums, track, res)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}
