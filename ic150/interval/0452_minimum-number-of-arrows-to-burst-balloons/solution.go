/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 11:01'
*/

package solution

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	res := 1
	pos := points[0][1]
	for i := 0; i < len(points); i++ {
		if points[i][0] <= pos {
			continue
		}

		res++
		pos = points[i][1]
	}

	return res
}
