/*
 * Author: robin-luo
 * Created time: 2024-02-19 16:46:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 16:51:52
 */

package solution

func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	var dfs func(path []int, steps int)
	dfs = func(path []int, steps int) {
		if steps == len(graph)-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(graph[steps]); i++ {
			next := append(path, graph[steps][i])
			dfs(next, graph[steps][i])
		}
	}

	dfs([]int{0}, 0)
	return res
}
