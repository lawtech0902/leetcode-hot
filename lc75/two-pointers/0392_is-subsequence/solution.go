/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-26 10:44:43
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 10:46:49
 * @Description:
 */

package solution

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}
