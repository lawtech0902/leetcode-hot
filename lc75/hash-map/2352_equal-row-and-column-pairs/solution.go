/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 10:37'
*/

package solution

import "fmt"

func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := make(map[string]int)
	for _, row := range grid {
		cnt[fmt.Sprint(row)]++
	}

	res := 0
	for j := 0; j < n; j++ {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, grid[i][j])
		}

		if val, ok := cnt[fmt.Sprint(arr)]; ok {
			res += val
		}
	}

	return res
}

func equalPairs1(grid [][]int) int {
	res := 0
	n := len(grid)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if equal(row, col, grid) {
				res++
			}
		}
	}

	return res
}

func equal(row, col int, grid [][]int) bool {
	n := len(grid)
	for i := 0; i < n; i++ {
		if grid[row][i] != grid[i][col] {
			return false
		}
	}

	return true
}
