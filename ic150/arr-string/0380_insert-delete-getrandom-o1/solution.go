/*
__author__ = 'robin-luo'
__date__ = '2023/12/06 9:49'
*/

package solution

import "math/rand"

type RandomizedSet struct {
	buff  []int
	cache map[int]int // num->idx
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		buff:  make([]int, 0),
		cache: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.cache[val]; ok {
		return false
	}

	this.buff = append(this.buff, val)
	this.cache[val] = len(this.buff) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.cache[val]; ok {
		this.cache[this.buff[len(this.buff)-1]] = idx
		this.buff[idx], this.buff[len(this.buff)-1] = this.buff[len(this.buff)-1], this.buff[idx]
		this.buff = this.buff[:len(this.buff)-1]
		delete(this.cache, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.buff) > 0 {
		return this.buff[rand.Intn(len(this.buff))]
	}

	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
