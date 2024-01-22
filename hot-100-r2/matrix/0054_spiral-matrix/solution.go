/*
__author__ = 'robin-luo'
__date__ = '2024/01/22 11:12'
*/

package solution

import "reflect"

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	for len(matrix) != 0 {
		res = append(res, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}

		matrix = turn(matrix)
	}

	return res
}

func turn(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	var newMatrix [][]int
	for i := 0; i < n; i++ {
		var subArr []int
		for j := 0; j < m; j++ {
			subArr = append(subArr, matrix[j][i])
		}

		newMatrix = append(newMatrix, subArr)
	}

	reverseAnySlice(newMatrix)
	return newMatrix
}

func reverseAnySlice(s any) {
	n := reflect.ValueOf(s).Len()
	swapper := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swapper(i, j)
	}
}

func spiralOrder1(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	u, d, l, r := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := l; i <= r; i++ {
			res = append(res, matrix[u][i])
		}

		u++
		if u > d {
			break
		}

		for i := u; i <= d; i++ {
			res = append(res, matrix[i][r])
		}

		r--
		if r < l {
			break
		}

		for i := r; i >= l; i-- {
			res = append(res, matrix[d][i])
		}

		d--
		if d < u {
			break
		}

		for i := d; i >= u; i-- {
			res = append(res, matrix[i][l])
		}

		l++
		if l > r {
			break
		}
	}

	return res
}
