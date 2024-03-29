/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 19:21'
*/

package solution

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	p, q, r, s := 0, 0, 1, 1
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = s
		s = p + q + r
	}

	return s
}
