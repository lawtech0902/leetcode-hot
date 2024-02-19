/*
 * Author: robin-luo
 * Created time: 2024-02-19 20:14:00
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 20:24:56
 */

package solution

func pacificAtlantic(heights [][]int) [][]int {
	var (
		m, n     = len(heights), len(heights[0])
		res      [][]int
		dfs      func(i, j int, visited [][]bool)
		pacific  = make([][]bool, m)
		atlantic = make([][]bool, m)
	)

	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	dfs = func(i, j int, visited [][]bool) {
		dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		visited[i][j] = true
		for k := 0; k < 4; k++ {
			x, y := i+dirs[k][0], j+dirs[k][1]
			if 0 <= x && x < m && 0 <= y && y < n && !visited[x][y] && heights[x][y] >= heights[i][j] {
				dfs(x, y, visited)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	for j := 0; j < n; j++ {
		dfs(0, j, pacific)
		dfs(m-1, j, atlantic)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
