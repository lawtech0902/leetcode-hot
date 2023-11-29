/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 15:17'
*/

package solution

func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 32; i++ {
		aBit, bBit, cBit := (a>>i)&1, (b>>i)&1, (c>>i)&1
		if cBit == 0 {
			res += aBit + bBit
		} else if aBit+bBit == 0 {
			res += 1
		}
	}

	return res
}
