/*
 * Author: robin-luo
 * Created time: 2024-02-02 19:43:54
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 20:01:27
 */

package solution

func wiggleMaxLength(nums []int) int {
	res := 1
	larger := 0
	for i := range nums {
		if larger == 0 {
			if nums[i] != nums[i-1] {
				res++
				if nums[i] > nums[i-1] {
					larger = 1
				} else {
					larger = -1
				}
			}
		} else if larger == 1 && nums[i] < nums[i-1] {
			res++
			larger = -1
		} else if larger == -1 && nums[i] > nums[i-1] {
			res++
			larger = 1
		}
	}

	return res
}
