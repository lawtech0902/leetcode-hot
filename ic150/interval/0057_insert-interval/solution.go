/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 10:48'
*/

package solution

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	left, right := newInterval[0], newInterval[1]
	merged := false
	for _, interval := range intervals {
		if interval[0] > right {
			if !merged {
				res = append(res, []int{left, right})
				merged = true
			}

			res = append(res, interval)
		} else if interval[1] < left {
			res = append(res, interval)
		} else {
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}

	if !merged {
		res = append(res, []int{left, right})
	}

	return
}
