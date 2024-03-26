/*
 * Author: robin-luo
 * Created time: 2024-03-18 16:53:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-18 16:55:20
 */

package solution

import "bytes"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	rows := make([]bytes.Buffer, numRows)
	for i := range rows {
		rows[i] = bytes.Buffer{}
	}

	i := 0
	flag := -1
	for _, ch := range s {
		rows[i].WriteRune(ch)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}

		i += flag
	}

	var res bytes.Buffer
	for _, row := range rows {
		res.WriteString(row.String())
	}

	return res.String()
}
