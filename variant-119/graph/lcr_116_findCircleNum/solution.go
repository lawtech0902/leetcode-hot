/*
 * Author: robin-luo
 * Created time: 2024-02-20 21:11:32
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 09:42:33
 */

package solution

// dfs
func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true
		for to, conn := range isConnected[from] {
			if !visited[to] && conn == 1 {
				dfs(to)
			}
		}
	}

	var res int
	for i, v := range visited {
		if !v {
			res++
			dfs(i)
		}
	}

	return res
}

// bfs
// bfs
func findCircleNum1(isConnected [][]int) int {
	var res int
	visited := make([]bool, len(isConnected))
	for i, v := range visited {
		if !v {
			res++
			q := []int{i}
			for len(q) > 0 {
				from := q[0]
				q = q[1:]
				visited[from] = true
				for to, conn := range isConnected[from] {
					if conn == 1 && !visited[to] {
						q = append(q, to)
					}
				}
			}
		}
	}

	return res
}
