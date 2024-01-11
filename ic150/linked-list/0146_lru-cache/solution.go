/*
__author__ = 'robin-luo'
__date__ = '2024/01/08 9:53'
*/

package solution

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	size, cap  int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:  0,
		cap:   capacity,
		cache: make(map[int]*Node),
		head:  &Node{},
		tail:  &Node{},
	}

	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head
	return lruCache
}

func (this *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) addToTail(node *Node) {
	this.tail.prev.next = node
	node.prev = this.tail.prev
	node.next = this.tail
	this.tail.prev = node
}

func (this *LRUCache) moveToTail(node *Node) {
	this.removeNode(node)
	this.addToTail(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		return node.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.moveToTail(node)
		node.value = value
		return
	}

	if this.size == this.cap {
		this.size--
		node := this.head.next
		this.removeNode(node)
		delete(this.cache, node.key)
	}

	node := &Node{key: key, value: value}
	this.size++
	this.addToTail(node)
	this.cache[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
