/*
 * Author: robin-luo
 * Created time: 2024-02-29 15:28:04
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 15:39:17
 */

package solution

import "strconv"

func addStrings(num1, num2 string) string {
	carry := 0
	res := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}

		if j >= 0 {
			y = int(num2[j] - '0')
		}

		sum := x + y + carry
		sum, carry = sum%10, sum/10
		res = strconv.Itoa(sum) + res
	}

	return res
}
