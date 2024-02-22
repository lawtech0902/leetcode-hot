/*
 * Author: robin-luo
 * Created time: 2024-02-20 15:56:27
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-20 16:02:31
 */

package solution

func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res [][]int
		dfs func(path []int, step int)
	)

	dfs = func(path []int, step int) {
		if step == len(graph)-1 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(graph[step]); i++ {
			next := append(path, graph[step][i])
			dfs(next, graph[step][i])
		}
	}

	dfs([]int{0}, 0)
	return res
}
