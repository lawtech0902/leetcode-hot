/*
 * Author: robin-luo
 * Created time: 2024-02-21 13:50:46
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 14:26:01
 */

package solution

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	dummy := &Node{
		Val:   0,
		Prev:  nil,
		Next:  root,
		Child: nil,
	}

	dfs(dummy, root)
	dummy.Next.Prev = nil
	return dummy.Next
}

func dfs(prev, cur *Node) *Node {
	if cur == nil {
		return prev
	}

	cur.Prev = prev
	prev.Next = cur
	tempNext := cur.Next
	tail := dfs(cur, cur.Child)
	cur.Child = nil
	return dfs(tail, tempNext)
}
