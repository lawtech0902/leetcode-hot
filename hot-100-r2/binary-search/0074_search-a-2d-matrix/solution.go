/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 15:31'
*/

package solution

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (l + r) / 2
		temp := matrix[mid/n][mid%n]
		if temp == target {
			return true
		} else if temp < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
