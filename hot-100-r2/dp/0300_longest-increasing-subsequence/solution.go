/*
__author__ = 'robin-luo'
__date__ = '2024/01/29 19:33'
*/

package solution

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	maxLen := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

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
