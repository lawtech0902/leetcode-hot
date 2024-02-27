/*
 * Author: robin-luo
 * Created time: 2024-02-26 19:56:14
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-26 20:25:10
 */

package main

import "fmt"

func getMaxNumLTN(nums []int, n int) int {
	var res int
	backTracking(nums, []int{}, n, 0, &res)
	return res
}

func backTracking(nums, track []int, target, sum int, res *int) {
	if sum >= target {
		return
	}

	*res = max(*res, sum)
	for i := 0; i < len(nums); i++ {
		track = append(track, nums[i])
		sum = sum*10 + nums[i]
		backTracking(nums, track, target, sum, res)
		sum = sum / 10
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(getMaxNumLTN([]int{1, 2, 4, 9}, 2533))
	fmt.Println(getMaxNumLTN([]int{1, 2, 4, 9}, 1111))
}
