/*
 * Author: robin-luo
 * Created time: 2024-02-03 18:32:58
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 18:58:25
 */

package solution

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}

		return people[i][0] > people[j][0]
	})

	res := make([][]int, 0)
	for _, p := range people {
		if p[1] >= len(res) {
			res = append(res, p)
		} else {
			res = append(res[:p[1]], append([][]int{p}, res[p[1]:]...)...)
		}
	}

	return res
}
