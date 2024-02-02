/*
 * Author: robin-luo
 * Created time: 2024-02-02 10:08:48
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 10:27:02
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
		track = append(track, s[startIndex:i+1])
		if i-startIndex+1 <= 3 && len(track) <= 4 && isValidIp(s, startIndex, i) {
			backtracking(s, i+1, track, res)
		} else {
			return
		}
		track = track[:len(track)-1]
	}
}

func isValidIp(s string, start, end int) bool {
	val, _ := strconv.Atoi(s[start : end+1])
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}

	if val > 255 {
		return false
	}

	return true
}
