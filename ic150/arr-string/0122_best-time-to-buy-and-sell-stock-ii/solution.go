/*
__author__ = 'robin-luo'
__date__ = '2023/12/05 16:48'
*/

package solution

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
