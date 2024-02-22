/*
 * Author: robin-luo
 * Created time: 2024-02-21 14:26:20
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 15:43:37
 */

package solution

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, insertVal int) *Node {
	node := &Node{Val: insertVal}
	if head == nil {
		node.Next = node
		return node
	}

	if head.Next == head {
		head.Next = node
		node.Next = head
		return head
	}

	cur, next := head, head.Next
	for next != head {
		if cur.Val <= insertVal && insertVal <= next.Val {
			break
		}

		if cur.Val > next.Val {
			if cur.Val < insertVal || insertVal < next.Val {
				break
			}
		}

		cur = cur.Next
		next = next.Next
	}

	cur.Next = node
	node.Next = next
	return head
}
