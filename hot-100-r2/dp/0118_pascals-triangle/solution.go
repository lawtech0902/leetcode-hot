/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 16:30'
*/

package solution

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := range res {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		res[i] = temp
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}
