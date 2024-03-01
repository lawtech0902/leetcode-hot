/*
 * Author: robin-luo
 * Created time: 2024-02-29 10:34:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 10:35:38
 */

package solution

import "math"

func maxProfit(prices []int) int {
	buy, sell := math.MaxInt64, 0
	for _, price := range prices {
		buy = min(buy, price)
		sell = max(sell, price-buy)
	}

	return sell
}
