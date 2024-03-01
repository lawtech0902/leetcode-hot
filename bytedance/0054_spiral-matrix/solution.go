/*
 * Author: robin-luo
 * Created time: 2024-02-29 11:12:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 11:33:10
 */

package solution

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
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
