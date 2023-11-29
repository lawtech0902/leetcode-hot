/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 10:43'
*/

package solution

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	res := make([]int, len(spells))
	for i, x := range spells {
		res[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}

	return res
}
