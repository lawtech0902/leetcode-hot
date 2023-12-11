/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 11:34'
*/

package solution

func canCompleteCircuit(gas, cost []int) int {
	totalTank, curTank := 0, 0
	startStation := 0
	for i := range gas {
		totalTank += gas[i] - cost[i]
		curTank += gas[i] - cost[i]
		if curTank < 0 {
			startStation = i + 1
			curTank = 0
		}
	}

	if totalTank >= 0 {
		return startStation
	} else {
		return -1
	}
}
