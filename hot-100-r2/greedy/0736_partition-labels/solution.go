/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 16:06'
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
