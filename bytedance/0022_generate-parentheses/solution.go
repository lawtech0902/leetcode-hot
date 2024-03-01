/*
 * Author: robin-luo
 * Created time: 2024-03-01 09:30:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 09:38:17
 */

package solution

func generateParenthesis(n int) []string {
	var res []string
	backtracking(n, 0, 0, "", &res)
	return res
}

func backtracking(n, l, r int, track string, res *[]string) {
	if n == l && n == r {
		*res = append(*res, track)
		return
	}

	if l < n {
		track += "("
		l++
		backtracking(n, l, r, track, res)
		l--
		track = track[:len(track)-1]
	}

	if l > r {
		track += ")"
		r++
		backtracking(n, l, r, track, res)
		r--
		track = track[:len(track)-1]
	}
}
