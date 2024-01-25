/*
__author__ = 'robin-luo'
__date__ = '2024/01/24 10:05'
*/

package solution

type Node struct {
	Val          int
	Next, Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return node
	}

	if n, ok := cachedNode[node]; ok {
		return n
	}

	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
