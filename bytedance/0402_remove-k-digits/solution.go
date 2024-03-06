/*
 * Author: robin-luo
 * Created time: 2024-03-05 10:15:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 10:52:41
 */

package solution

import "strings"

// https://leetcode.cn/problems/remove-k-digits/solutions/290203/yi-zhao-chi-bian-li-kou-si-dao-ti-ma-ma-zai-ye-b-5/
func removeKdigits(num string, k int) string {
	var stack []rune
	remain := len(num) - k
	for _, digit := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > digit {
			stack = stack[:len(stack)-1]
			k--
		}

		stack = append(stack, digit)
	}

	res := strings.TrimLeft(string(stack[:remain]), "0")
	if res == "" {
		return "0"
	}

	return res
}
