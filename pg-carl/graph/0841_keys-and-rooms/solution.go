/*
 * Author: robin-luo
 * Created time: 2024-02-20 11:26:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 11:36:15
 */

package solution

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	visited := make([]bool, n)
	var dfs func(key int)
	dfs = func(key int) {
		if visited[key] {
			return
		}

		visited[key] = true
		keys := rooms[key]
		for _, k := range keys {
			dfs(k)
		}
	}

	dfs(0)
	for _, v := range visited {
		if !v {
			return false
		}
	}

	return true
}
