/*
 * Author: robin-luo
 * Created time: 2024-03-13 10:08:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 10:14:07
 */

package solution

func intToRoman(num int) string {
	ints := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := range ints {
		for num >= ints[i] {
			num -= ints[i]
			res += romans[i]
		}
	}

	return res
}
