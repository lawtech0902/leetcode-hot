/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 16:15:42
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 16:26:15
 * @Description:
 */

package solution

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	bytes := []byte(s)
	for l < r {
		if !isVowel(bytes[l]) {
			l++
			continue
		}

		if !isVowel(bytes[r]) {
			r--
			continue
		}

		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	for _, item := range []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'} {
		if b == item {
			return true
		}
	}

	return false
}
