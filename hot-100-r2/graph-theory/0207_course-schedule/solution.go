/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 15:29'
*/

package solution

// bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges = make([][]int, numCourses)
		indeg = make([]int, numCourses)
		res   []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

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

	return len(res) == numCourses
}
