/*
 * Author: robin-luo
 * Created time: 2024-03-05 14:19:14
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 14:27:16
 */

package solution

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.InStack = append(this.InStack, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.OutStack) == 0 {
		if len(this.InStack) == 0 {
			return -1
		}

		for len(this.InStack) > 0 {
			this.OutStack = append(this.OutStack, this.InStack[len(this.InStack)-1])
			this.InStack = this.InStack[:len(this.InStack)-1]
		}
	}

	v := this.OutStack[len(this.OutStack)-1]
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return v
}
