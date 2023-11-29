/*
__author__ = 'robin-luo'
__date__ = '2023/11/29 09:47'
*/

package solution

import "math"

type StockSpanner struct {
	stack [][2]int
	idx   int
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: [][2]int{{-1, math.MaxInt32}},
		idx:   -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	this.idx++
	for price >= this.stack[len(this.stack)-1][1] {
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, [2]int{this.idx, price})
	return this.idx - this.stack[len(this.stack)-2][0]
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
