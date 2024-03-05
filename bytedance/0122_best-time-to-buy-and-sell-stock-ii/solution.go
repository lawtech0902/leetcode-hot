/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:06:17
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 11:08:48
 */

package solution

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}
