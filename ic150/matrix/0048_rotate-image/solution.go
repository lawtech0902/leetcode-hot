/*
__author__ = 'robin-luo'
__date__ = '2023/12/20 11:42'
*/

package solution

func rotate(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := i; j < size; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < size; i++ {
		reverseSlice(matrix[i])
	}
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
