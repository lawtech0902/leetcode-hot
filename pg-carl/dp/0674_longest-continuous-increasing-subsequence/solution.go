/*
 * Author: robin-luo
 * Created time: 2024-02-19 13:45:58
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 13:52:46
 */

package solution

func findLengthOfLCIS(nums []int) int {
	res, start := 0, 0
	for i := range nums {
		if i >= 1 && nums[i] <= nums[i-1] {
			start = i
		}

		res = max(res, i-start+1)
	}

	return res
}
