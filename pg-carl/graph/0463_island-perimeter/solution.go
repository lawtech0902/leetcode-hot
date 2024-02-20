/*
 * Author: robin-luo
 * Created time: 2024-02-20 11:47:23
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 13:42:45
 */

package solution

func islandPerimeter(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				// 只要遇到一对相邻陆地，总周长-2
				// up存在相邻陆地
				if i-1 >= 0 && grid[i-1][j] == 1 {
					res -= 2
				}

				// left存在相邻陆地
				if j-1 >= 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res
}
