/*
 * Author: robin-luo
 * Created time: 2024-01-30 20:19:43
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-30 20:25:00
 */

package solution

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}

	return slow
}
