/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:12:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 11:14:10
 */

package solution

// 展开
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) >> 1
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}
