/*
__author__ = 'robin-luo'
__date__ = '2024/01/16 9:33'
*/

package solution

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := hash[target-num]; ok {
			return []int{i, j}
		}

		hash[num] = i
	}

	return nil
}
