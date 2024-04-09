/*
 * Author: robin-luo
 * Created time: 2024-03-29 15:14:55
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-31 19:03:35
 */

package solution

type Cell struct {
	h, x, y int
}

type MinHeap []Cell

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].h < mh[j].h
}

// 最小堆
func trapRainWater(heightMap [][]int) int {
	var res int
	m, n := len(heightMap), len(heightMap[0])
	if m <= 2 || n <= 2 {
		return res
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	mh := &MinHeap{}
}
