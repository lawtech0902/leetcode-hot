/*
__author__ = 'robin-luo'
__date__ = '2023/12/20 10:53'
*/

package solution

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]int)
	cols := make(map[int]map[byte]int)
	boxes := make(map[int]map[byte]int)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num != '.' {
				boxIndex := (i/3)*3 + j/3
				if _, ok := rows[i]; !ok {
					rows[i] = make(map[byte]int)
				}

				if _, ok := cols[j]; !ok {
					cols[j] = make(map[byte]int)
				}

				if _, ok := boxes[boxIndex]; !ok {
					boxes[boxIndex] = make(map[byte]int)
				}

				if _, ok := rows[i][num]; ok {
					return false
				} else {
					rows[i][num] = 1
				}

				if _, ok := cols[j][num]; ok {
					return false
				} else {
					cols[j][num] = 1
				}

				if _, ok := boxes[boxIndex][num]; ok {
					return false
				} else {
					boxes[boxIndex][num] = 1
				}
			}
		}
	}

	return true
}
