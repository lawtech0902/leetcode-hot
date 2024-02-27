/*
 * Author: robin-luo
 * Created time: 2024-02-27 16:06:31
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 16:09:52
 */

package solution

import "sort"

func merge(intervals [][]int) [][]int {
	n := len(intervals)
	if n == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < n; i++ {
		if intervals[i][0] <= res[len(res)-1][1] {
			res[len(res)-1][1] = max(res[len(res)-1][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}
