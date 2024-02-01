/*
 * Author: robin-luo
 * Created time: 2024-02-01 20:01:19
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-01 20:12:07
 */

package solution

var digitsArr = [10]string{
	"",     // 0
	"",     // 1
	"abc",  // 2
	"def",  // 3
	"ghi",  // 4
	"jkl",  // 5
	"mno",  // 6
	"pqrs", // 7
	"tuv",  // 8
	"wxyz", // 9
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 || n > 4 {
		return nil
	}

	var res []string
	backtracking(digits, "", 0, &res)
	return res
}

func backtracking(digits, track string, index int, res *[]string) {
	if len(track) == len(digits) {
		*res = append(*res, track)
		return
	}

	letters := digitsArr[digits[index]-'0']
	for i := 0; i < len(letters); i++ {
		track += string(letters[i])
		backtracking(digits, track, index+1, res)
		track = track[:len(track)-1]
	}
}
