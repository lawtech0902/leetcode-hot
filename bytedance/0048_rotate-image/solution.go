/*
 * Author: robin-luo
 * Created time: 2024-03-04 14:30:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 14:47:14
 */

package solution

func rotate(matrix [][]int) {
	n := len(matrix)
	// 水平对折翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
