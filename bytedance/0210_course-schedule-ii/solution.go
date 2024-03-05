/*
 * Author: robin-luo
 * Created time: 2024-03-04 20:36:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 20:39:58
 */

package solution

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indeg := make([]int, numCourses)
	courses := make([]int, 0)
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

	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		courses = append(courses, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(courses) == numCourses {
		return courses
	}

	return nil
}
