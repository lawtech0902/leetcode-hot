/*
 * Author: robin-luo
 * Created time: 2024-02-28 09:21:58
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-28 09:34:08
 */

package solution

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Size  int
	Cap   int
	Cache map[int]*Node
	Head  *Node
	Tail  *Node
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		Size:  0,
		Cap:   capacity,
		Cache: make(map[int]*Node),
		Head:  &Node{},
		Tail:  &Node{},
	}

	lc.Head.Next = lc.Tail
	lc.Tail.Prev = lc.Head
	return lc
}

func (this *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	node.Next = nil
	node.Prev = nil
}

func (this *LRUCache) addToTail(node *Node) {
	this.Tail.Prev.Next = node
	node.Prev = this.Tail.Prev
	node.Next = this.Tail
	this.Tail.Prev = node
}

func (this *LRUCache) moveToTail(node *Node) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveToTail(node)
		return node.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		this.moveToTail(node)
		node.Value = value
		return
	}

	if this.Size == this.Cap {
		this.Size--
		node := this.Head.Next
		this.removeNode(node)
		delete(this.Cache, key)
	}

	node := &Node{
		Key:   key,
		Value: value,
	}

	this.Size++
	this.addToTail(node)
	this.Cache[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
