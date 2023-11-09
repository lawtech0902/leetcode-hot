/*
__author__ = 'robin-luo'
__date__ = '2023/11/09 10:24'
*/

package solution

var (
	num int
	vis []bool
)

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	num = 0
	vis = make([]bool, n)
	dfs(rooms, 0)
	return num == n
}

func dfs(rooms [][]int, x int) {
	vis[x] = true
	num++
	for _, it := range rooms[x] {
		if !vis[it] {
			dfs(rooms, it)
		}
	}
}
