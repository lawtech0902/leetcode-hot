/*
__author__ = 'robin-luo'
__date__ = '2023/12/21 10:03'
*/

package solution

func setZeroes(matrix [][]int) {
	rowNum := len(matrix)
	colNum := len(matrix[0])
	row := make([]bool, rowNum)
	col := make([]bool, colNum)

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}

	for i := 0; i < rowNum; i++ {
		for j := 0; j < colNum; j++ {
			if row[i] || col[j] {
				matrix[i][j] = 0
			}
		}
	}
}
