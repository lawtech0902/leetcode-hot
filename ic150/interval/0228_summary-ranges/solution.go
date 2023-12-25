/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 10:05'
*/

package solution

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var (
		res  []string
		head int
	)

	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		if head == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			temp := fmt.Sprintf("%d->%d", nums[head], nums[i])
			res = append(res, temp)
		}

		head = i + 1
	}

	return res
}
