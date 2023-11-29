/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-11-17 11:07:21
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-11-21 11:06:43
 * @Description:
 */

package solution

type Node struct {
	dest   string
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string][]Node)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		value := values[i]
		graph[a] = append(graph[a], Node{dest: b, weight: value})
		graph[b] = append(graph[b], Node{dest: a, weight: 1 / value})
	}

	res := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		_, ok1 := graph[start]
		_, ok2 := graph[end]
		if !ok1 || !ok2 {
			res[i] = -1.0
		} else {
			visited := make(map[string]bool)
			res[i] = dfs(graph, start, end, 1.0, visited)
		}
	}

	return res
}

func dfs(graph map[string][]Node, curNode, endNode string, curRes float64, visited map[string]bool) float64 {
	if curNode == endNode {
		return curRes
	}

	visited[curNode] = true
	for _, node := range graph[curNode] {
		if !visited[node.dest] {
			res := dfs(graph, node.dest, endNode, curRes*node.weight, visited)
			if res != -1 {
				return res
			}
		}
	}

	return -1.0
}
