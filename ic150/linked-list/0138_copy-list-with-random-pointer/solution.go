/*
__author__ = 'robin-luo'
__date__ = '2023/12/26 14:55'
*/

package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
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

func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{
			Val:  node.Val,
			Next: node.Next,
		}
	}

	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}

	newHead := head.Next
	for node := head; node != nil; node = node.Next {
		newNode := node.Next
		node.Next = node.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
	}

	return newHead
}
