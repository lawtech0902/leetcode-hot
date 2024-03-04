/*
 * Author: robin-luo
 * Created time: 2024-03-01 10:43:34
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 11:00:36
 */

package solution

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	iMin, iMax := 0, m
	halfLen := (m + n + 1) / 2

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i

		if i < m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > 0 && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
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
