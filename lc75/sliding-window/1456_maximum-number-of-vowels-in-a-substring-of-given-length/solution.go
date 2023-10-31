/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-27 09:47:44
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-27 10:33:27
 * @Description:
 */

package solution

import "math"

func maxVowels(s string, k int) int {
	size := len(s)
	vowelCount := 0
	for i := 0; i < k; i++ {
		vowelCount += isVowel(s[i])
	}

	res := vowelCount
	for i := k; i < size; i++ {
		vowelCount += isVowel(s[i]) - isVowel(s[i-k])
		res = int(math.Max(float64(res), float64(vowelCount)))
	}

	return res
}

func isVowel(ch byte) int {
	switch ch {
	case 'a', 'e', 'i', 'o', 'u':
		return 1
	default:
		return 0
	}
}
