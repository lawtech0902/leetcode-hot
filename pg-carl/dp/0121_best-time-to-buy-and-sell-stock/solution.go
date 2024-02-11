/*
 * Author: robin-luo
 * Created time: 2024-02-08 16:16:35
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 16:19:54
 */

package solution

func maxProfit(prices []int) int {
	buy, sell := 1<<63-1, 0
	for _, price := range prices {
		buy = min(buy, price)
		sell = max(sell, price-buy)
	}

	return sell
}
