/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 15:27:14
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 16:15:13
 * @Description:
 */

package solution

func canPlaceFlowers(flowerbed []int, n int) bool {
	// 前后补0解决边界问题
	flowerbed = append([]int{0}, flowerbed...)
	flowerbed = append(flowerbed, 0)
	count := 0
	i := 1
	length := len(flowerbed)
	for i < length-1 {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			count++
			i += 2
		} else {
			i++
		}
	}

	return count >= n
}
