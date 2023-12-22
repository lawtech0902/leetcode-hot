/*
__author__ = 'robin-luo'
__date__ = '2023/12/22 14:57'
*/

package solution

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{i, j}
		}

		m[num] = i
	}

	return nil
}
