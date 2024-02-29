/*
 * Author: robin-luo
 * Created time: 2024-02-28 10:22:06
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 11:43:57
 */

package solution

// dfs
func canMeasureWater(x, y, t int) bool {
	var dfs func(int, int) bool
	memo := make(map[[2]int]bool)
	dfs = func(remainX, remainY int) bool {
		if remainX+remainY == t {
			return true
		}

		if memo[[2]int{remainX, remainY}] {
			return false
		}

		memo[[2]int{remainX, remainY}] = true
		// 尝试所有操作
		// 倒满任意一个容器
		if dfs(x, remainY) || dfs(remainX, y) {
			return true
		}
		// 倒空任意一个容器
		if dfs(0, remainY) || dfs(remainX, 0) {
			return true
		}
		// 将 x 倒入 y
		if dfs(max(remainX-(y-remainY), 0), min(remainX+remainY, y)) {
			return true
		}
		// 将 y 倒入 x
		if dfs(min(x, remainX+remainY), max(0, remainY-(x-remainX))) {
			return true
		}

		return false
	}

	return dfs(0, 0)
}

// math
// 贝祖定理
// ax+by = target, 则target必为x和y的最大公约数的倍数
func canMeasureWater1(x, y, t int) bool {
	if x+y < t {
		return false
	}

	return t%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
