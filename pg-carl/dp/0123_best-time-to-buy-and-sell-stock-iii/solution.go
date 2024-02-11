/*
 * Author: robin-luo
 * Created time: 2024-02-08 16:33:38
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-11 15:17:20
 */

package solution

import "math"

func maxProfit(prices []int) int {
	buy1, sell1 := math.MaxInt64, 0
	buy2, sell2 := math.MinInt64, 0
	for _, price := range prices {
		buy1 = min(buy1, price)
		sell1 = max(sell1, price-buy1)
		buy2 = max(buy2, sell1-price)
		sell2 = max(sell2, buy2+price)
	}

	return sell2
}
