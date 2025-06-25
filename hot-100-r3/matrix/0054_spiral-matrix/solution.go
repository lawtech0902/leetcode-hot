/*
__author__ = 'robin-luo'
__date__ = '2025/06/17 15:16'
*/

package solution

func spiralOrder(matrix [][]int) []int {
	res := make([]int, 0)
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			res = append(res, matrix[up][i])
		}

		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			res = append(res, matrix[i][right])
		}

		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			res = append(res, matrix[down][i])
		}

		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			res = append(res, matrix[i][left])
		}

		left++
		if left > right {
			break
		}
	}

	return res
}
