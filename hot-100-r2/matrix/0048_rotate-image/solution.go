/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 14:23'
*/

package solution

func rotate(matrix [][]int) {
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		reverseSlice(matrix[i])
	}
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
