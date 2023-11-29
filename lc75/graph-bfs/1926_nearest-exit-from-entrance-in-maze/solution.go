/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-11-21 11:09:13
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-11-21 19:39:10
 * @Description:
 */

package solution

type Point struct {
	x, y  int
	steps int
}

func nearestExit(maze [][]byte, entrance []int) int {
	directions := [][]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}
	queue := []Point{Point{entrance[0], entrance[1], 0}}
	visited := make([][]bool, len(maze))
	for i := range visited {
		visited[i] = make([]bool, len(maze[0]))
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if (curr.x != entrance[0] || curr.y != entrance[1]) && (curr.x == 0 || curr.x == len(maze)-1 || curr.y == 0 || curr.y == len(maze[0])-1) {
			return curr.steps
		}

		for _, d := range directions {
			nx := curr.x + d[0]
			ny := curr.y + d[1]
			if nx >= 0 && nx < len(maze) && ny >= 0 && ny < len(maze[0]) && maze[nx][ny] == '.' && !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, Point{nx, ny, p.steps + 1})
			}
		}
	}

	return -1
}
