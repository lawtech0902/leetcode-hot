/*
 * Author: robin-luo
 * Created time: 2024-03-04 19:46:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-17 20:33:20
 */

package main

import "strings"

func multiply(num1, num2 string) string {
	l1, l2 := len(num1), len(num2)
	byteRes := make([]byte, l1+l2+1)
	for i := range byteRes {
		byteRes[i] = '0'
	}

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			temp := (num1[i] - '0') * (num2[j] - '0')
			pos := i + j + 2
			byteRes[pos-1] += (temp + byteRes[pos] - '0') / 10
			byteRes[pos] = (temp+byteRes[pos]-'0')%10 + '0'
		}
	}

	res := strings.TrimLeft(string(byteRes), "0")
	if res == "" {
		return "0"
	}

	return res
}
