/*
 * Author: robin-luo
 * Created time: 2024-02-02 14:10:09
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-02 14:42:55
 */

package solution

import (
	"sort"
)

type Pair struct {
	Target  string
	Visited bool
}

type Pairs []*Pair

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Less(i, j int) bool {
	return p[i].Target < p[j].Target
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func findItinerary(tickets [][]string) []string {
	var res []string
	res = append(res, "JFK")

	// map[出发机场]pair(目的地，是否被访问过)
	targets := make(map[string]Pairs)
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil {
			targets[ticket[0]] = make(Pairs, 0)
		}

		targets[ticket[0]] = append(targets[ticket[0]], &Pair{
			Target:  ticket[1],
			Visited: false,
		})
	}

	for _, v := range targets {
		sort.Sort(v)
	}

	backtracking(tickets, targets, &res)
	return res
}

func backtracking(tickets [][]string, targets map[string]Pairs, res *[]string) bool {
	if len(tickets)+1 == len(*res) {
		return true
	}

	for _, pair := range targets[(*res)[len(*res)-1]] {
		if !pair.Visited {
			*res = append(*res, pair.Target)
			pair.Visited = true
			if backtracking(tickets, targets, res) {
				return true
			}

			pair.Visited = false
			*res = (*res)[:len(*res)-1]
		}
	}

	return false
}
