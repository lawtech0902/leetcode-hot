/*
__author__ = 'robin-luo'
__date__ = '2023/12/25 9:36'
*/

package solution

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			if i-m[v] <= k {
				return true
			}
		}

		m[v] = i
	}

	return false
}
