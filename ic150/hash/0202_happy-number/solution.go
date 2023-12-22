/*
__author__ = 'robin-luo'
__date__ = '2023/12/22 15:58'
*/

package solution

func isHappy(n int) bool {
	slow, fast := n, n
	for {
		fast = squareSum(fast)
		slow = squareSum(slow)
		slow = squareSum(slow)
		if slow == fast {
			break
		}
	}

	return fast == 1
}

func squareSum(n int) int {
	sum := 0
	for n != 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}

	return sum
}
