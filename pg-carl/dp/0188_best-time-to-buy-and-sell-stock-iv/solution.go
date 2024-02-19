/*
 * Author: robin-luo
 * Created time: 2024-02-11 15:17:41
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-11 15:24:58
 */

package solution

import "math"

func maxProfit(k int, prices []int) int {
	k = min(k, len(prices)/2) + 1
	buy := make([]int, k)
	sell := make([]int, k)
	for i := range buy {
		buy[i] = math.MinInt64
	}

	for _, price := range prices {
		for j := 1; j < k; j++ {
			buy[j] = max(buy[j], sell[j-1]-price)
			sell[j] = max(sell[j], buy[j]+price)
		}
	}

	return sell[k-1]
}
