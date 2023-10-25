/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 15:09:49
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 15:23:20
 * @Description:
 */

package solution

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	res := make([]bool, n)
	for i, candy := range candies {
		res[i] = candy+extraCandies >= maxCandies
	}

	return res
}
