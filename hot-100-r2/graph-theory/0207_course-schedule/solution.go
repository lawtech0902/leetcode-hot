/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 15:29'
*/

package solution

// bfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 初始化邻接表和入度数组
	var (
		edges = make([][]int, numCourses) // 邻接表：edges[i]表示课程 i 的所有后续课程
		indeg = make([]int, numCourses)   // 入度数组：indeg[i]表示课程 i 的先修课程数量
		res   []int                       // 存储拓扑排序结果
	)

	// 构建图结构
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0]) // info[1] -> info[0]
		indeg[info[0]]++                                 // info[0] 的入度++
	}

	// 初始化队列：将所有入度为 0 的课程加入队列
	var q []int
	for i := 0; i < numCourses; i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	// 拓扑排序过程
	for len(q) > 0 {
		u := q[0] // 取出队列中的第一个课程
		q = q[1:]
		res = append(res, u)

		// 遍历该课程的所有后续课程
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return len(res) == numCourses
}
