/*
 * Author: robin-luo
 * Created time: 2024-02-19 11:14:40
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-19 11:40:40
 */

package solution

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	res := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		res = max(res, dp[i])
	}

	return res
}

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	tail := 1
	d := make([]int, n+1)
	d[tail] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > d[tail] {
			tail++
			d[tail] = nums[i]
		} else {
			l, r := 1, tail
			pos := 0
			for l <= r {
				mid := (l + r) >> 1
				if d[mid] < nums[i] {
					pos = mid
					l = mid + 1
				} else {
					r = mid - 1
				}
			}

			d[pos+1] = nums[i]
		}
	}

	return tail
}
