/*
 * Author: robin-luo
 * Created time: 2024-02-20 16:02:53
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 16:37:01
 */

package solution

type Node struct {
	Dest   string
	Weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Node)
	for i, e := range equations {
		a, b := e[0], e[1]
		value := values[i]
		graph[a] = append(graph[a], Node{Dest: b, Weight: value})
		graph[b] = append(graph[b], Node{Dest: a, Weight: 1 / value})
	}

	var dfs func(cur, end string, curRes float64, visited map[string]bool) float64
	dfs = func(cur, end string, curRes float64, visited map[string]bool) float64 {
		if cur == end {
			return curRes
		}

		visited[cur] = true
		for _, node := range graph[cur] {
			if !visited[node.Dest] {
				res := dfs(node.Dest, end, curRes*node.Weight, visited)
				if res != -1.0 {
					return res
				}
			}
		}

		return -1.0
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]
		_, ok1 := graph[start]
		_, ok2 := graph[end]
		if !ok1 || !ok2 {
			res[i] = -1.0
		} else {
			visited := make(map[string]bool)
			res[i] = dfs(start, end, 1.0, visited)
		}
	}

	return res
}
