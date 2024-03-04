/*
 * Author: robin-luo
 * Created time: 2024-03-01 14:10:12
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-01 14:26:08
 */

package solution

// 双栈法
type MyQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.InStack = append(this.InStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutStack) == 0 {
		for len(this.InStack) != 0 {
			this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
			this.InStack = this.InStack[:len(this.InStack)-1]
		}
	}

	v := this.OutStack[len(this.OutStack)-1]
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return v
}

func (this *MyQueue) Peek() int {
	if len(this.OutStack) == 0 {
		for len(this.InStack) != 0 {
			this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
			this.InStack = this.InStack[:len(this.InStack)-1]
		}
	}

	return this.OutStack[len(this.OutStack)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.InStack) == 0 && len(this.OutStack) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
