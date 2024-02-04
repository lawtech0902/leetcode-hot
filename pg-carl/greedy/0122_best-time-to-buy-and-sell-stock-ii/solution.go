/*
 * Author: robin-luo
 * Created time: 2024-02-02 21:07:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 12:54:21
 */

package solution

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
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
