/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 10:28'
*/

package solution

func generateParenthesis(n int) []string {
	var res []string
	backtracking(n, 0, 0, "", &res)
	return res
}

func backtracking(n, nLeft, nRight int, temp string, res *[]string) {
	if nLeft == n && nRight == n {
		*res = append(*res, temp)
		return
	}

	if nLeft < n {
		temp += "("
		nLeft++
		backtracking(n, nLeft, nRight, temp, res)
		nLeft--
		temp = temp[:len(temp)-1]
	}

	if nLeft > nRight {
		temp += ")"
		nRight++
		backtracking(n, nLeft, nRight, temp, res)
		nRight--
		temp = temp[:len(temp)-1]
	}
}
