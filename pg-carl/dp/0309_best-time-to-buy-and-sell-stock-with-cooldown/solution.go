/*
 * Author: robin-luo
 * Created time: 2024-02-11 15:25:28
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 11:03:38
 */

package solution

import "math"

/*
sold：
前一天hold，当日卖出股票

hold： 可由两个情况转换来
前一天hold，当日rest
前一天rest，当日买入股票变为hold

rest：
前一天sold，当日必须rest
前一天rest，当日继续rest
*/
func maxProfit(prices []int) int {
	sold, rest, hold := 0, 0, math.MinInt32
	for _, p := range prices {
		preSold := sold
		sold = hold + p
		hold = max(hold, rest-p)
		rest = max(rest, preSold)
	}

	return max(sold, rest)
}
