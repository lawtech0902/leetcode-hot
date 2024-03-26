/*
 * Author: robin-luo
 * Created time: 2024-03-19 11:20:57
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 11:22:33
 */

package solution

func sortArrayByParityII(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 1 {
			for nums[j]%2 == 1 {
				j += 2
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	return nums
}
