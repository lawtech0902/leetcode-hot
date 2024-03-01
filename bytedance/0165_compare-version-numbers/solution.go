/*
 * Author: robin-luo
 * Created time: 2024-02-29 17:03:53
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 17:07:58
 */

package solution

func compareVersion(v1, v2 string) int {
	i, j := 0, 0
	for i < len(v1) || j < len(v2) {
		a, b := 0, 0
		for i < len(v1) && v1[i] != '.' {
			a = int(v1[i]-'0') + a*10
			i++
		}

		for j < len(v2) && v2[j] != '.' {
			b = int(v2[j]-'0') + b*10
			j++
		}

		if a > b {
			return 1
		} else if a < b {
			return -1
		}

		i++
		j++
	}

	return 0
}
