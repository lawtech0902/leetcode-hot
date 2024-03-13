/*
 * Author: robin-luo
 * Created time: 2024-03-13 15:17:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 15:36:48
 */

package solution

import (
	"slices"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return abs(arr[i]-x) < abs(arr[j]-x)
	})

	arr = arr[:k]
	slices.Sort(arr)
	return arr
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
