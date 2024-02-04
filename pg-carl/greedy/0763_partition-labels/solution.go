/*
 * Author: robin-luo
 * Created time: 2024-02-04 09:45:43
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-04 10:01:22
 */

package solution

func partitionLabels(s string) []int {
	lastPos := [26]int{}
	for i, ch := range s {
		lastPos[ch-'a'] = i
	}

	var res []int
	start, end := 0, 0
	for i, ch := range s {
		if lastPos[ch-'a'] > end {
			end = lastPos[ch-'a']
		}

		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}

	return res
}
