/*
 * Author: robin-luo
 * Created time: 2024-02-02 16:54:26
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 19:42:21
 */

package solution

import "slices"

func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	i, j := 0, 0
	res := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res++
			i++
		}
		j++
	}
	return res
}
