/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 10:27'
*/

package solution

import "sort"

type Interval struct {
	Start, End int
}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	its := make([]Interval, 0)
	for _, it := range intervals {
		its = append(its, Interval{
			Start: it[0],
			End:   it[1],
		})
	}

	sort.Slice(its, func(i, j int) bool {
		return its[i].Start < its[j].Start
	})

	var (
		newIts []Interval
		res    [][]int
	)

	temp := its[0]
	for i := 1; i < len(its); i++ {
		if its[i].Start <= temp.End {
			temp.End = max(temp.End, its[i].End)
		} else {
			newIts = append(newIts, temp)
			temp = its[i]
		}
	}

	newIts = append(newIts, temp)
	for _, newIt := range newIts {
		res = append(res, []int{newIt.Start, newIt.End})
	}

	return res
}
