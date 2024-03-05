/*
 * Author: robin-luo
 * Created time: 2024-03-05 09:23:31
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 09:31:12
 */

package main

import "fmt"

// 循环依赖检测。如，[['A', 'B'], ['B', 'C'], ['C', 'D'], ['B', 'D']] => false，[['A', 'B'], ['B', 'C'], ['C', 'A']] => true（2021.4 字节跳动-幸福里-后端）[2]
// 手撕代码：小王写了一个makefile，其中有n个编译项编号为0~n-1，他们互相之间有依赖关系。请写一个程序解析依赖，给出一个可行的编译顺序。（2021.03 字节跳动-系统部-后端）[3]

// 现有n个编译项，编号为0 ~ n-1。给定一个二维数组，表示编译项之间有依赖关系。如[0, 1]表示1依赖于0。
// 若存在循环依赖则返回空；不存在依赖则返回可行的编译顺序。

func haveCircularDependency(n int, prerequisites [][]int) bool {
	edges := make([][]int, n)
	indeg := make([]int, n)
	res := make([]int, 0)
	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}

	var q []int
	for i, v := range indeg {
		if v == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		u := q[0]
		q = q[1:]
		res = append(res, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return len(res) == n
}

func main() {
	fmt.Println(haveCircularDependency(5, [][]int{
		{0, 2},
		{1, 2},
		{2, 3},
		{2, 4},
	}))
}
