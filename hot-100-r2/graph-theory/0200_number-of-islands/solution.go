/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 13:39'
*/

package solution

// dfs
func numIslands(grid [][]byte) int {
	var (
		res int
		dfs func(i, j int)
	)

	dfs = func(i, j int) {
		if 0 <= i && i < len(grid) && 0 <= j && j < len(grid[i]) && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i+1, j)
			dfs(i-1, j)
			dfs(i, j+1)
			dfs(i, j-1)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

type UnionFind struct {
	parent []int
	count  int
}

func NewUnionFind(grid [][]byte) *UnionFind {
	m, n := len(grid), len(grid[0])
	count := 0
	parent := make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parent[i*n+j] = i*n + j
				count++
			}
		}
	}

	return &UnionFind{
		parent: parent,
		count:  count,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}

	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) {
	rootX, rootY := uf.Find(x), uf.Find(y)
	if rootX != rootY {
		uf.parent[rootX] = rootY
		uf.count--
	}
}

// union-find
func numIslands1(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	uf := NewUnionFind(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				for _, dir := range directions {
					newX, newY := i+dir[0], i+dir[1]
					if newX >= 0 && newX < m && newY >= 0 && newY < n && grid[newX][newY] == '1' {
						uf.Union(i*n+j, newX*n+newY)
					}
				}
			}
		}
	}

	return uf.count
}
