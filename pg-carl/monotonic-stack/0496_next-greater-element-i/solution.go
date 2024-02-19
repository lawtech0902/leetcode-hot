/*
 * Author: robin-luo
 * Created time: 2024-02-19 15:25:16
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 15:33:25
 */

package solution

func nextGreaterElement(nums1, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	res := make([]int, m)
	for i := range res {
		res[i] = -1
	}

	hash := make(map[int]int, m)
	for i, v := range nums1 {
		hash[v] = i
	}

	stack := []int{0}
	for i := 1; i < n; i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			index2 := stack[len(stack)-1]
			if index1, ok := hash[nums2[index2]]; ok {
				res[index1] = nums2[i]
			}

			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	return res
}
