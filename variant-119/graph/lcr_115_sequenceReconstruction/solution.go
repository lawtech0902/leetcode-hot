/*
 * Author: robin-luo
 * Created time: 2024-02-20 20:55:59
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 21:10:59
 */

package solution

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	edges := make([][]int, n+1)
	indeg := make([]int, n+1)
	for _, seq := range sequences {
		for i := 1; i < len(seq); i++ {
			x, y := seq[i-1], seq[i]
			edges[x] = append(edges[x], y)
			indeg[y]++
		}
	}

	var q []int
	for i, v := range indeg {
		if i >= 1 && v == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}

		u := q[0]
		q = q[1:]
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return true
}
