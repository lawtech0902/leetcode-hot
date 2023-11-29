/*
__author__ = 'robin-luo'
__date__ = '2023/11/27 10:31'
*/

package solution

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var guess func(num int) int

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}
