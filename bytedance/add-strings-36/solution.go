/*
 * Author: robin-luo
 * Created time: 2024-03-05 15:59:10
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 16:29:26
 */

// 36进制由0-9，a-z，共36个字符表示。
// 要求按照加法规则计算出任意两个36进制正整数的和，如1b + 2x = 48  （解释：47+105=152）
// 要求：不允许使用先将36进制数字整体转为10进制，相加后再转回为36进制的做法

package main

import (
	"fmt"
)

func add36Strings(num1 string, num2 string) string {
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	var res string
	for i >= 0 || j >= 0 || carry > 0 {
		x, y := 0, 0
		if i >= 0 {
			x = getInt(num1[i])
		}

		if j >= 0 {
			y = getInt(num2[j])
		}

		temp := x + y + carry
		res = getChar(temp%36) + res
		carry = temp / 36
		i--
		j--
	}

	return res
}

func getChar(n int) string {
	if n <= 9 {
		return string(n + '0')
	} else {
		return string(n - 10 + 'a')
	}
}

func getInt(ch byte) int {
	if '0' <= ch && ch <= '9' {
		return int(ch - '0')
	} else {
		return int(ch - 'a' + 10)
	}
}

func main() {
	a, b := "1b", "2x"
	c := add36Strings(a, b)
	fmt.Println(c)
}
