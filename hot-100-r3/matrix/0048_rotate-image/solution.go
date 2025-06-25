/*
__author__ = 'robin-luo'
__date__ = '2025/06/18 09:46'
*/

package solution

func rotate(matrix [][]int) {
	n := len(matrix)
	// 先转置，行列交换
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 再每行反转
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
