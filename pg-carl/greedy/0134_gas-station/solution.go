/*
 * Author: robin-luo
 * Created time: 2024-02-03 17:18:56
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-03 17:37:17
 */

package solution

func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	totalSum := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 {
			start = i + 1
			curSum = 0
		}
	}

	if totalSum < 0 {
		return -1
	}

	return start
}
