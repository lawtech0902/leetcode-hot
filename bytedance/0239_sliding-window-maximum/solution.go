/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:12:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 17:41:03
 */

package solution

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, 0, n-k+1)
	deque := make([]int, 0)
	for i, v := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= v {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i-deque[0] >= k {
			deque = deque[1:]
		}

		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
