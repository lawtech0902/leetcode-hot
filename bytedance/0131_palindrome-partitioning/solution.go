/*
 * Author: robin-luo
 * Created time: 2024-03-13 09:21:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 09:36:16
 */

package solution

func partition(s string) [][]string {
	var (
		track []string
		res   [][]string
	)

	backtracking(s, 0, track, &res)
	return res
}

func backtracking(s string, startIndex int, track []string, res *[][]string) {
	if startIndex == len(s) {
		temp := make([]string, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			track = append(track, s[startIndex:i+1])
			backtracking(s, i+1, track, res)
			track = track[:len(track)-1]
		}
	}
}

func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}

		start++
		end--
	}

	return true
}
