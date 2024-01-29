/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 14:41'
*/

package solution

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	low := prices[0]
	res := 0
	for _, price := range prices {
		if price < low {
			low = price
		}

		res = max(res, price-low)
	}

	return res
}
