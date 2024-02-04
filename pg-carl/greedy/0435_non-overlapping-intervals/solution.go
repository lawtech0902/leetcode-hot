/*
 * Author: robin-luo
 * Created time: 2024-02-03 21:15:21
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 21:42:27
 */

package solution

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			intervals[i][1] = min(intervals[i][1], intervals[i-1][1])
			count++
		}
	}

	return count
}
