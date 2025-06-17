/*
__author__ = 'robin-luo'
__date__ = '2025/06/16 15:00'
*/

package solution

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0, len(nums)-k+1)
	dq := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 移除过期队头元素
		if len(dq) > 0 && dq[0] < i-k+1 {
			dq = dq[1:]
		}

		// 维护单调性，从队尾移除小于当前值的元素
		for len(dq) > 0 && nums[dq[len(dq)-1]] < nums[i] {
			dq = dq[:len(dq)-1]
		}

		// 当前索引入队
		dq = append(dq, i)
		// 如果当前已经足够覆盖窗口，则队头为当前窗口最大值
		if i+1 >= k {
			res = append(res, nums[dq[0]])
		}
	}

	return res
}
