/*
 * Author: robin-luo
 * Created time: 2024-02-27 15:59:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 09:42:37
 */

package solution

import "math/rand"

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

func sortArray1(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	pivotIndex := rand.Intn(n)
	nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]
	pivot := nums[0]
	i, j := -1, n
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	quickSort(nums[:j+1])
	quickSort(nums[j+1:])
}
