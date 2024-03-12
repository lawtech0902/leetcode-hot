/*
__author__ = 'robin-luo'
__date__ = '2024/01/17 10:03'
*/

package solution

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	res := make([]int, 0, n-k+1)
	deque := make([]int, 0, k)
	for i := 0; i < n; i++ {
		for len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)
		if i-k+1 >= 0 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
