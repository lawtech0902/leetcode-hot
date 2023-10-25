/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-07-27 13:55:35
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-07-27 13:55:46
 * @Description:
 */

package solution

func mergeAlternately(word1, word2 string) string {
	var res string
	len1 := len(word1)
	len2 := len(word2)
	for i := 0; i < min(len1, len2); i++ {
		res += string(word1[i])
		res += string(word2[i])
	}

	if len1 > len2 {
		res += word1[len2:]
	}

	if len2 > len1 {
		res += word2[len1:]
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
