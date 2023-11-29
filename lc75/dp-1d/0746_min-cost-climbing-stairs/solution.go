/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 19:38'
*/

package solution

func minCostClimbingStairs(cost []int) int {
	a, b := 0, 0
	n := len(cost)
	c := 0
	for i := 2; i < n+1; i++ {
		c = min(a+cost[i-2], b+cost[i-1])
		a = b
		b = c
	}

	return c
}
