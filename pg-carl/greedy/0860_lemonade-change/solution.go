/*
 * Author: robin-luo
 * Created time: 2024-02-03 18:13:21
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 18:31:10
 */

package solution

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five < 0 {
				return false
			}

			ten++
			five--
		} else {
			if five > 0 && ten > 0 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}

	return true
}
