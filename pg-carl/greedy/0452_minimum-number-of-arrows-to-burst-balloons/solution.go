/*
 * Author: robin-luo
 * Created time: 2024-02-03 19:18:37
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 20:55:25
 */

package solution

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] > points[i-1][1] {
			res++
		} else {
			points[i][1] = min(points[i-1][1], points[i][1])
		}
	}

	return res
}
