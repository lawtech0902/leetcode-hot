/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 21:18'
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

func backtracking(digits, temp string, index int, res *[]string) {
	if len(temp) == len(digits) {
		*res = append(*res, temp)
		return
	}

	tempIndex := digits[index] - '0'
	letters := digitsArr[tempIndex]
	for i := 0; i < len(letters); i++ {
		temp += string(letters[i])
		backtracking(digits, temp, index+1, res)
		temp = temp[:len(temp)-1]
	}
}
