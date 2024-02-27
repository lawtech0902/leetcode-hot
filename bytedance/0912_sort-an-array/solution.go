/*
 * Author: robin-luo
 * Created time: 2024-02-27 15:59:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 16:04:12
 */

package solution

// merge sort
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var res []int
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
