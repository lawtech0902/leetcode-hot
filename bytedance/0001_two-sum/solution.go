/*
 * Author: robin-luo
 * Created time: 2024-02-28 09:35:13
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 09:36:21
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
