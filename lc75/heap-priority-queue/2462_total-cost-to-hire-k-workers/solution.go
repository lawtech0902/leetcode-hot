/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 09:29'
*/

package solution

import (
	"container/heap"
	"sort"
)

type MinHeap struct {
	sort.IntSlice
}

func (h MinHeap) Push(any) {

}

func (h MinHeap) Pop() (_ any) {
	return
}

func totalCost(costs []int, k, candidates int) int64 {
	res := 0
	if n := len(costs); candidates*2 < n {
		pre := MinHeap{costs[:candidates]}
		heap.Init(&pre)
		suf := MinHeap{costs[n-candidates:]}
		heap.Init(&suf)
		for i, j := candidates, n-1-candidates; k > 0 && i <= j; k-- {
			if pre.IntSlice[0] <= suf.IntSlice[0] {
				res += pre.IntSlice[0]
				pre.IntSlice[0] = costs[i]
				heap.Fix(&pre, 0)
				i++
			} else {
				res += suf.IntSlice[0]
				suf.IntSlice[0] = costs[j]
				heap.Fix(&suf, 0)
				j--
			}
		}

		costs = append(pre.IntSlice, suf.IntSlice...)
	}

	sort.Ints(costs)
	for _, c := range costs[:k] {
		res += c
	}

	return int64(res)
}
