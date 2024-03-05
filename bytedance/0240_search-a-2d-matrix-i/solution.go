/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:09:13
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 11:11:12
 */

package solution

// 右上角开始二分
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			row++
		} else {
			col--
		}
	}

	return false
}
