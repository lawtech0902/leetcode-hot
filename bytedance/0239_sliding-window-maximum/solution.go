/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:12:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 10:59:26
 */

package solution

func maxSlidingWindow(nums []int, k int) []int {
	// 获取数组长度
	n := len(nums)
	// 初始化一个切片，用于存放结果，长度为 n-k+1
	res := make([]int, 0, n-k+1)
	// 初始化一个队列，用于存放索引
	deque := make([]int, 0)
	// 遍历数组
	for i, v := range nums {
		// 遍历队列，将队列中元素小于当前元素的元素出队
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= v {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素索引入队
		deque = append(deque, i)
		// 如果队列首元素索引与当前元素索引差值大于等于 k，将队列首元素出队
		for i-deque[0] >= k {
			deque = deque[1:]
		}

		// 如果当前元素索引大于等于 k-1，将队列首元素添加到结果切片中
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	// 返回结果切片
	return res
}
