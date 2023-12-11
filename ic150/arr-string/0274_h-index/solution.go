/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 17:24'
*/

package solution

import "sort"

func hIndex(citations []int) (h int) {
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0 && citations[i] > h; i-- {
		h++
	}

	return
}

func hIndex1(citations []int) int {
	left, right := 0, len(citations)
	var mid int
	for left < right {
		mid = (left + right + 1) >> 1
		cnt := 0
		for _, v := range citations {
			if v >= mid {
				cnt++
			}
		}

		if cnt >= mid {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return left
}
