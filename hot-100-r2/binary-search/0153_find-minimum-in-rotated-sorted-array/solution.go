/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 16:32'
*/

package solution

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}

	return nums[low]
}
