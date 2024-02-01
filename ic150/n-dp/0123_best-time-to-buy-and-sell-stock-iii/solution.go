/*
 * Author: robin-luo
 * Created time: 2024-01-31 14:58:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-01-31 15:10:04
 */

package solution

import "math"

func maxProfit(prices []int) int {
	buy1, sell1 := math.MinInt64, 0
	buy2, sell2 := math.MinInt64, 0
	for _, price := range prices {
		buy1 = max(buy1, -price)
		sell1 = max(sell1, buy1+price)
		buy2 = max(buy2, sell1-price)
		sell2 = max(sell2, buy2+price)
	}

	return sell2
}
