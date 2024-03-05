/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:16:44
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 11:27:35
 */

package solution

type MinStack struct {
	Vals []int
	Mins []int
}

func Constructor() MinStack {
	return MinStack{
		Vals: make([]int, 0),
		Mins: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Vals = append(this.Vals, val)
	if len(this.Mins) == 0 {
		this.Mins = append(this.Mins, val)
	} else {
		if this.Mins[len(this.Mins)-1] > val {
			this.Mins = append(this.Mins, val)
		} else {
			this.Mins = append(this.Mins, this.Mins[len(this.Mins)-1])
		}
	}
}

func (this *MinStack) Pop() {
	this.Vals = this.Vals[:len(this.Vals)-1]
	this.Mins = this.Mins[:len(this.Mins)-1]
}

func (this *MinStack) Top() int {
	return this.Vals[len(this.Vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.Mins[len(this.Mins)-1]
}
