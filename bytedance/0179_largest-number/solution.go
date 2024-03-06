/*
 * Author: robin-luo
 * Created time: 2024-03-05 16:32:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 16:37:53
 */

package solution

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	n := len(nums)
	numStrArr := make([]string, n)
	for i, num := range nums {
		numStrArr[i] = strconv.Itoa(num)
	}

	sort.Slice(numStrArr, func(i, j int) bool {
		if numStrArr[i][0] == numStrArr[j][0] {
			return numStrArr[i]+numStrArr[j] > numStrArr[j]+numStrArr[i]
		}

		return numStrArr[i] > numStrArr[j]
	})

	res := strings.Join(numStrArr, "")
	if res[0] == '0' {
		return "0"
	}

	return res
}
