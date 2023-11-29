/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 11:10'
*/

package solution

func maxProfit(prices []int, fee int) int {
	notHold, hold := 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		notHold = max(notHold, hold+prices[i]-fee)
		hold = max(hold, notHold-prices[i])
	}

	return notHold
}
