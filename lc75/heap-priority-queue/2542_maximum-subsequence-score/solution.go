/*
__author__ = 'robin-luo'
__date__ = '2023/11/23 11:47'
*/

package solution

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	*h = old[:n-1]
	return old[n-1]
}

type Pair struct {
	x, y int
}

func maxScore(nums1, nums2 []int, k int) int64 {
	var nums []Pair
	for i, num1 := range nums1 {
		nums = append(nums, Pair{num1, nums2[i]})
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i].y > nums[j].y
	})

	h := MinHeap{}
	currSum := 0
	for i := 0; i < k; i++ {
		h = append(h, nums[i].x)
		currSum += nums[i].x
	}

	heap.Init(&h)
	res := currSum * nums[k-1].y
	for _, p := range nums[k:] {
		if p.x > h[0] {
			currSum = currSum - h[0] + p.x
			heap.Pop(&h)
			heap.Push(&h, p.x)
			if t := currSum * p.y; t > res {
				res = t
			}
		}
	}

	return int64(res)
}
