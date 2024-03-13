/*
 * Author: robin-luo
 * Created time: 2024-03-13 10:00:15
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 10:04:33
 */

package solution

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]
	for _, str := range strs[1:] {
		for strings.Index(str, prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}
