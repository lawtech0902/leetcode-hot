/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-11-21 19:40:29
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-11-22 10:51:02
 * @Description:
 */

package solution

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	queue := make([][]int, 0)
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			} else if grid[i][j] == 1 {
				count++
			}
		}
	}

	if count == 0 {
		return 0
	}

	minutes := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			row, col := curr[0], curr[1]
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					queue = append(queue, []int{newRow, newCol})
					count--
				}
			}
		}

		if len(queue) > 0 {
			minutes++
		}
	}

	if count > 0 {
		return -1
	}

	return minutes
}
