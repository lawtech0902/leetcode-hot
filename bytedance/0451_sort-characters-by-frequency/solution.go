/*
 * Author: robin-luo
 * Created time: 2024-03-19 15:07:29
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 15:10:23
 */

package solution

import (
	"strings"
)

func frequencySort(s string) string {
	bucket := make([][]rune, len(s)+1)
	memo := make(map[rune]int)
	var res string
	for _, r := range s {
		memo[r]++
	}

	for r, cnt := range memo {
		bucket[cnt] = append(bucket[cnt], r)
	}

	for i := len(s); i >= 0; i-- {
		for _, r := range bucket[i] {
			res += strings.Repeat(string(r), i)
		}
	}

	return res
}
