/*
 * Author: robin-luo
 * Created time: 2024-02-03 13:21:53
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 16:51:33
 */

package solution

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	curDistance := 0  // 当前覆盖最远距离下标
	steps := 0        // 记录走的最大步数
	nextDistance := 0 // 下一步覆盖最远距离下标
	for i := 0; i < len(nums)-1; i++ {
		nextDistance = max(nextDistance, i+nums[i])
		if i == curDistance {
			curDistance = nextDistance
			steps++
		}
	}

	return steps
}
