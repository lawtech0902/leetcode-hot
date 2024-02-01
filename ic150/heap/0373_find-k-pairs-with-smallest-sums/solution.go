/*
 * Author: robin-luo
 * Created time: 2024-02-01 13:50:02
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 15:54:07
 */

package solution

import "container/heap"

type Tuple struct {
	sum, i, j int
}

type MinHeap []Tuple

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].sum < mh[j].sum
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MinHeap) Push(v any) {
	*mh = append(*mh, v.(Tuple))
}

func (mh *MinHeap) Pop() any {
	v := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return v
}

func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	res := make([][]int, 0, min(k, m*n))
	mh := &MinHeap{}
	heap.Init(mh)
	heapSize := min(k, m)
	for i := 0; i < heapSize; i++ {
		heap.Push(mh, Tuple{nums1[i] + nums2[0], i, 0})
	}

	for len(*mh) > 0 && len(res) < k {
		p := heap.Pop(mh).(Tuple)
		i, j := p.i, p.j
		res = append(res, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(mh, Tuple{nums1[i] + nums2[j+1], i, j + 1})
		}
	}

	return res
}
