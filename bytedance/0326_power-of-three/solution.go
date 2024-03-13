/*
 * Author: robin-luo
 * Created time: 2024-03-12 21:56:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-13 09:20:50
 */

package solution

import "math"

// 判断是否为最大3的幂的约数
func isPowerOfThree(n int) bool {
	return n > 0 && int(math.Pow(3, 19))%n == 0
}
