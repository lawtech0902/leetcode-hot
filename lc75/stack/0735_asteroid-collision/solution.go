/*
__author__ = 'robin-luo'
__date__ = '2023/10/31 11:27'
*/

package solution

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, v := range asteroids {
		alive := true
		for alive && v < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = stack[len(stack)-1] < -v
			if stack[len(stack)-1] <= -v {
				stack = stack[:len(stack)-1]
			}
		}

		if alive {
			stack = append(stack, v)
		}
	}

	return stack
}
