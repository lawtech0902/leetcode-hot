/*
 * Author: robin-luo
 * Created time: 2024-01-31 16:44:54
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 13:49:31
 */

package solution

import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x any) {
	*mh = append(*mh, x.(int))
}

func (mh *MaxHeap) Pop() any {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}

type Pair struct {
	c, p int
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	pairArr := make([]Pair, n)
	for i, profit := range profits {
		pairArr[i] = Pair{capital[i], profit}
	}

	sort.Slice(pairArr, func(i, j int) bool {
		return pairArr[i].c < pairArr[j].c
	})

	mh := &MaxHeap{}
	for cur := 0; k > 0; k-- {
		for cur < n && pairArr[cur].c <= w {
			heap.Push(mh, pairArr[cur].p)
			cur++
		}

		if mh.Len() == 0 {
			break
		}

		w += heap.Pop(mh).(int)
	}

	return w
}
