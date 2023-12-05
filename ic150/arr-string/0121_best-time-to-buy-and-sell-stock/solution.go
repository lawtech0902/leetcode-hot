/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 16:24'
*/

package solution

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < low {
			low = price
		}

		maxProfit = max(maxProfit, price-low)
	}

	return maxProfit
}
