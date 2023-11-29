/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 11:40'
*/

package solution

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := range res {
		res[i] = onesCount(i)
	}

	return res
}

func onesCount(x int) (ones int) {
	for ; x > 0; x &= x - 1 {
		ones++
	}

	return ones
}

func countBits1(n int) []int {
	res := make([]int, n+1)
	highBit := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			highBit = i
		}

		res[i] = res[i-highBit] + 1
	}

	return res
}
