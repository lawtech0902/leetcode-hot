/*
__author__ = 'robin-luo'
__date__ = '2023/12/11 11:18'
*/

package solution

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var (
		cs    []uint8
		count int
		index int
	)

	for i := 0; i < numRows; i++ {
		count = 0
		index = count*(2*numRows-2) + i
		for index < len(s) {
			cs = append(cs, s[index])
			for i != 0 && i != numRows-1 {
				index = (count+1)*(2*numRows-2) - i
				if index >= len(s) {
					break
				}

				cs = append(cs, s[index])
			}

			count++
			index = count*(2*numRows-2) + i
		}
	}

	return string(cs)
}
