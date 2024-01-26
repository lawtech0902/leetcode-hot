/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 16:01'
*/

package solution

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			res := []int{0, 0}
			if nums[l] == target {
				res[0] = l
			}

			if nums[r] == target {
				res[1] = r
			}

			for i := mid; i < r+1; i++ {
				if nums[i] != target {
					res[1] = i - 1
					break
				}
			}

			for i := mid; i > l-1; i-- {
				if nums[l] != target {
					res[0] = i + 1
					break
				}
			}

			return res
		}
	}

	return []int{-1, -1}
}
