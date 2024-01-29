/*
__author__ = 'robin-luo'
__date__ = '2024/01/28 19:22'
*/

package solution

type MinStack struct {
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		vals: make([]int, 0),
		mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.vals = append(this.vals, val)
	lenMins := len(this.mins)
	if lenMins == 0 {
		this.mins = append(this.mins, val)
	} else {
		if this.mins[lenMins-1] > val {
			this.mins = append(this.mins, val)
		} else {
			this.mins = append(this.mins, this.mins[lenMins-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.vals = this.vals[:len(this.vals)-1]
	this.mins = this.mins[:len(this.mins)-1]
}

func (this *MinStack) Top() int {
	return this.vals[len(this.vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
