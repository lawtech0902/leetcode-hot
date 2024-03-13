/*
 * Author: robin-luo
 * Created time: 2024-03-12 21:28:51
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-12 21:55:33
 */

package main

import (
	"fmt"
	"strconv"
)

func numsSameConsecDiff(n, k int) []int {
	var (
		track []byte
		res   []int
	)

	backtracking(n, k, track, &res)
	return res
}

func backtracking(n, k int, track []byte, res *[]int) {
	if len(track) == n {
		v, _ := strconv.Atoi(string(track))
		*res = append(*res, v)
		return
	}

	if len(track) > n {
		return
	}

	for i := 0; i <= 9; i++ {
		if len(track) == 0 && i == 0 {
			continue
		}

		if len(track) == 0 {
			track = append(track, byte('0'+i))
			backtracking(n, k, track, res)
			track = track[:len(track)-1]
		} else {
			if abs(i-int(track[len(track)-1]-'0')) == k {
				track = append(track, byte('0'+i))
				backtracking(n, k, track, res)
				track = track[:len(track)-1]
			}
		}
	}
}

func abs(n int) int {
	if n > 0 {
		return n
	}

	return -n
}

func main() {
	fmt.Println(numsSameConsecDiff(3, 7))
}
