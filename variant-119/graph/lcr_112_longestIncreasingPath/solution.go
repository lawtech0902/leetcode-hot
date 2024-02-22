/*
 * Author: robin-luo
 * Created time: 2024-02-20 16:37:32
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 16:53:10
 */

package solution

func longestIncreasingPath(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := range cache {
		cache[i] = make([]int, n)
	}

	var dfs func(i, j, prev int) int
	dfs = func(i, j, prev int) int {
		if 0 <= i && i < m && 0 <= j && j < n {
			cur := matrix[i][j]
			if cur > prev {
				if cache[i][j] > 0 {
					return cache[i][j]
				}

				l := 1 + max(dfs(i-1, j, cur), dfs(i+1, j, cur), dfs(i, j-1, cur), dfs(i, j+1, cur))
				cache[i][j] = l
				return l
			}
		}

		return 0
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res = max(res, dfs(i, j, -1))
		}
	}

	return res
}
