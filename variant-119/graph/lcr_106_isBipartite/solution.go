/*
 * Author: robin-luo
 * Created time: 2024-02-20 13:58:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 14:22:26
 */

package solution

func isBipartite(graph [][]int) bool {
	var (
		UNCOLORED, RED, GREEN = 0, 1, 2
		n                     = len(graph)
		colors                = make([]int, n)
		valid                 = true
		dfs                   func(node, c int)
	)

	dfs = func(node, color int) {
		colors[node] = color
		colorNeigh := RED
		if color == RED {
			colorNeigh = GREEN
		}

		for _, neigh := range graph[node] {
			if colors[neigh] == UNCOLORED {
				dfs(neigh, colorNeigh)
				if !valid {
					return
				}
			} else if colors[neigh] != colorNeigh {
				valid = false
				return
			}
		}
	}

	for i := 0; i < n && valid; i++ {
		if colors[i] == UNCOLORED {
			dfs(i, RED)
		}
	}

	return valid
}
