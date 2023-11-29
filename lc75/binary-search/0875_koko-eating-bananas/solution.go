/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 11:18'
*/

package solution

import "sort"

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	return 1 + sort.Search(max-1, func(speed int) bool {
		speed++
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}

		return time <= h
	})
}
