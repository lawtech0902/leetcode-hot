/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-30 19:55:39
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-30 20:12:13
 * @Description:
 */

package solution

func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	sum := 0
	for i, v := range nums {
		if 2*sum+v == total {
			return i
		}

		sum += v
	}

	return -1
}
