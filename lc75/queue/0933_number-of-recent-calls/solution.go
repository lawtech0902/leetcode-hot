/*
__author__ = 'robin-luo'
__date__ = '2023/11/01 09:31'
*/

package solution

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.queue) > 0 && this.queue[0] < t-3000 {
		this.queue = this.queue[1:]
	}

	this.queue = append(this.queue, t)
	return len(this.queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
