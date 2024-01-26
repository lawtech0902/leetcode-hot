/*
__author__ = 'robin-luo'
__date__ = '2024/01/26 16:50'
*/

package solution

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	// 确保 nums1 是较短的数组
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	// 二分查找的边界，在 nums1 中搜索中位数的位置
	iMin, iMax := 0, m
	// 两个数组合并后中位数的位置，即 (m+n+1)/2。这里使用 (m+n+1) 是为了处理数组长度和为奇数的情况，保证中位数位于左半部分
	halfLen := (m + n + 1) / 2
	for iMin <= iMax {
		// i+j = halfLen 满足合并后的数组左半部分的长度
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < m && nums2[j-1] > nums1[i] {
			// 如果 i 小于 m 且 nums2[j-1] > nums1[i]，说明 i 的值太小，需要增大 i，因此将 iMin 更新为 i + 1
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			// 如果 i 大于 0 且 nums1[i-1] > nums2[j]，说明 i 的值太大，需要减小 i，因此将 iMax 更新为 i - 1
			iMax = i - 1
		} else {
			// nums1[i-1] <= nums2[j] 和 nums2[j-1] <= nums1[i]，即确保左半部分的最大值小于等于右半部分的最小值
			// 计算左半部分的最大值 maxOfLeft 和右半部分的最小值 minOfRight
			// 如果数组长度和为奇数，中位数即为左半部分的最大值
			// 如果数组长度和为偶数，中位数为左半部分的最大值和右半部分的最小值的平均值
			maxOfLeft := 0
			if i == 0 {
				maxOfLeft = nums2[j-1]
			} else if j == 0 {
				maxOfLeft = nums1[i-1]
			} else {
				maxOfLeft = max(nums1[i-1], nums2[j-1])
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			minOfRight := 0
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = min(nums1[i], nums2[j])
			}

			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}
