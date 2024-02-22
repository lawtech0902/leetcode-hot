/*
 * Author: robin-luo
 * Created time: 2024-02-20 17:00:48
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 20:39:59
 */

package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	var q []int
	for i, v := range indeg {
		if v == 0 {
			q = append(q, i)
		}
	}

	var res []int
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		res = append(res, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(res) == numCourses {
		return res
	} else {
		return nil
	}
}
