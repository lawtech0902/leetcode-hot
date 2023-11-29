/*
__author__ = 'robin-luo'
__date__ = '2023/11/15 11:32'
*/

package solution

func minReorder(n int, connections [][]int) int {
	count := 0
	graph := buildGraph(n, connections)
	visited := make([]bool, n)
	dfs(graph, visited, 0, &count)
	return count
}

func buildGraph(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}

	for _, connection := range connections {
		from, to := connection[0], connection[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], -from)
	}

	return graph
}

func dfs(graph [][]int, visited []bool, city int, count *int) {
	visited[city] = true
	for _, neighbour := range graph[city] {
		if !visited[abs(neighbour)] {
			if neighbour > 0 {
				*count++
			}

			dfs(graph, visited, abs(neighbour), count)
		}
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
