/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 11:01'
*/

package solution

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row := make([]bool, m)
	col := make([]bool, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
