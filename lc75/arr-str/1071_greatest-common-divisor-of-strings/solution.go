/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 14:02:01
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 15:03:10
 * @Description:
 */

package solution

import "strings"

func gcdOfStrings(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	gcdLen := gcd(len1, len2)
	subStr := str1[:gcdLen]
	if strings.Repeat(subStr, len1/gcdLen) == str1 && strings.Repeat(subStr, len2/gcdLen) == str2 {
		return subStr
	}

	return ""
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
