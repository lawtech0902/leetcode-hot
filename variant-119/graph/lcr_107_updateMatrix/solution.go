/*
 * Author: robin-luo
 * Created time: 2024-02-20 14:29:39
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 14:43:19
 */

package solution

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	q := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				matrix[i][j] = -1
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		i, j := q[0][0], q[0][1]
		q = q[1:]
		for _, v := range dirs {
			x, y := i+v[0], i+v[1]
			if 0 <= x && x < m && 0 <= y && y < n && matrix[x][y] == -1 {
				matrix[x][y] = matrix[i][j] + 1
				q = append(q, []int{x, y})
			}
		}
	}

	return matrix
}
