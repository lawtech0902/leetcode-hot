/*
 * Author: robin-luo
 * Created time: 2024-03-01 10:16:49
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 17:25:48
 */

package solution

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	var (
		track []string
		res   []string
	)

	backtracking(s, 0, track, &res)
	return res
}

func backtracking(s string, startIndex int, track []string, res *[]string) {
	if startIndex == len(s) && len(track) == 4 {
		temp := strings.Join(track, ".")
		*res = append(*res, temp)
	}

	for i := startIndex; i < len(s); i++ {
		if i-startIndex < 4 && len(track) <= 4 && isValidIP(s, startIndex, i) {
			track = append(track, s[startIndex:i+1])
			backtracking(s, i+1, track, res)
			track = track[:len(track)-1]
		} else {
			return
		}
	}
}

func isValidIP(s string, start, end int) bool {
	v, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}

	if v > 255 {
		return false
	}

	return true
}
