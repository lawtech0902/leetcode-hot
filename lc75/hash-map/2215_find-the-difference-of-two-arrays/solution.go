/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-30 20:13:02
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-31 09:52:15
 * @Description:
 */

package solution

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for _, v := range nums1 {
		map1[v] = true
	}

	for _, v := range nums2 {
		map2[v] = true
	}

	var diff1, diff2 []int
	for v := range map1 {
		if !map2[v] {
			diff1 = append(diff1, v)
		}
	}

	for v := range map2 {
		if !map1[v] {
			diff2 = append(diff2, v)
		}
	}

	return [][]int{diff1, diff2}
}
