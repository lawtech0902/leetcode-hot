/*
__author__ = 'robin-luo'
__date__ = '2023/11/15 09:23'
*/

package solution

// dfs
func findCircleNum(isConnected [][]int) int {
	var res int
	visited := make([]bool, len(isConnected))
	var dfs func(from int)
	dfs = func(from int) {
		visited[from] = true
		for to, conn := range isConnected[from] {
			if conn == 1 && !visited[to] {
				dfs(to)
			}
		}
	}

	for i, v := range visited {
		if !v {
			res++
			dfs(i)
		}
	}

	return res
}

// bfs
func findCircleNum1(isConnected [][]int) int {
	var res int
	visited := make([]bool, len(isConnected))
	for i, v := range visited {
		if !v {
			res++
			queue := []int{i}
			for len(queue) > 0 {
				from := queue[0]
				queue = queue[1:]
				visited[from] = true
				for to, conn := range isConnected[from] {
					if conn == 1 && !visited[to] {
						queue = append(queue, to)
					}
				}
			}
		}
	}

	return res
}
