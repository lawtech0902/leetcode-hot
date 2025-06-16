/*
__author__ = 'robin-luo'
__date__ = '2025/06/13 11:38'
*/

package solution

func trap(height []int) int {
	res := 0
	stack := []int{0} // 存柱子索引，初始放入第一个柱子的索引

	for i := 1; i < len(height); i++ { // 从第二个柱子开始遍历
		// 当前柱子高度 > 栈顶柱子高度，说明可能形成凹陷
		for len(stack) > 0 && stack[len(stack)-1] < height[i] {
			mid := height[stack[len(stack)-1]] // 栈顶柱子的高度
			stack = stack[:len(stack)-1]       // 栈顶柱子弹出
			if len(stack) > 0 {
				left := height[stack[len(stack)-1]]
				right := height[i]
				h := min(left, right) - mid
				w := i - stack[len(stack)-1] - 1
				res += h * w
			}
		}

		stack = append(stack, i)
	}

	return res
}
