/*
 * Author: robin-luo
 * Created time: 2024-02-19 20:30:41
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 10:41:51
 */

package solution

func largestIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	areaMap := make(map[int]int)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	mark := 2
	isAllGrid := true
	area := 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if visited[i][j] || grid[i][j] == 0 {
			return
		}

		visited[i][j] = true
		grid[i][j] = mark
		area++
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k][0], j+dirs[k][1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}

			dfs(x, y)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				isAllGrid = false
			}

			if !visited[i][j] && grid[i][j] == 1 {
				area = 0
				dfs(i, j)
				areaMap[mark] = area
				mark++
			}
		}
	}

	if isAllGrid {
		return m * n
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			area = 1
			if grid[i][j] == 0 {
				visitedGrid := make(map[int]bool)
				for k := 0; k < 4; k++ {
					x, y := i+dirs[k][0], j+dirs[k][1]
					if x < 0 || x >= m || y < 0 || y >= n || visitedGrid[grid[x][y]] {
						continue
					}

					mark = grid[x][y]
					area += areaMap[mark]
					visitedGrid[mark] = true
				}

				res = max(res, area)
			}
		}
	}

	return res
}
