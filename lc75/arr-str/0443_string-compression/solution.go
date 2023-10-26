/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-26 09:40:55
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-26 10:27:41
 * @Description:
 */

package solution

func compress(chars []byte) int {
	write, left := 0, 0
	for read, ch := range chars {
		if read == len(chars)-1 || ch != chars[read+1] {
			chars[write] = ch
			write++
			num := read - left + 1
			if num > 1 {
				anchor := write
				for ; num > 0; num /= 10 {
					chars[write] = '0' + byte(num%10)
					write++
				}

				s := chars[anchor:write]
				for i, n := 0, len(s); i < n/2; i++ {
					s[i], s[n-1-i] = s[n-1-i], s[i]
				}
			}

			left = read + 1
		}
	}

	return write
}
