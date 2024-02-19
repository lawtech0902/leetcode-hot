/*
 * Author: robin-luo
 * Created time: 2024-02-19 11:04:08
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 11:14:00
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
