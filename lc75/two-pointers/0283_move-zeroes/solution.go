/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-26 10:37:07
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 10:38:25
 * @Description:
 */

package solution

func moveZeroes(nums []int) {
	count := 0
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}

	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}
