/*
 * @Author: lawtech0902 584563542@qq.com
 * @Date: 2023-10-30 11:03:06
 * @LastEditors: lawtech0902 584563542@qq.com
 * @LastEditTime: 2023-10-30 19:54:43
 * @Description:
 */

package solution

func largestAltitude(gain []int) int {
	altitude, maxAltitude := 0, 0
	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude
}
