/*
__author__ = 'robin-luo'
__date__ = '2024/01/09 10:54'
*/

package solution

type Node struct {
	Val               int
	Left, Right, Next *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	cur := root
	for cur != nil {
		dummy := &Node{}
		pre := dummy
		for cur != nil {
			if cur.Left != nil {
				pre.Next = cur.Left
				pre = pre.Next
			}

			if cur.Right != nil {
				pre.Next = cur.Right
				pre = pre.Next
			}

			cur = cur.Next
		}

		cur = dummy.Next
	}

	return root
}
