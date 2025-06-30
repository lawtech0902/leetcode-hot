/*
__author__ = 'robin-luo'
__date__ = '2025/06/28 16:12'
*/

package solution

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var m = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 检查是否已经复制过这个节点
	if res := m[head]; res != nil {
		return res
	}

	// 创建新节点，复制
	res := &Node{Val: head.Val}
	// 在映射中记录原始节点到复制节点的关系
	m[head] = res
	// 递归复制Next和Random指针
	res.Next = copyRandomList(head.Next)
	res.Random = copyRandomList(head.Random)
	return res
}
