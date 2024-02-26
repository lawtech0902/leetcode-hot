/*
 * Author: robin-luo
 * Created time: 2024-02-22 10:19:19
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 10:19:41
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

	q := []*Node{root}
	for len(q) > 0 {
		temp := q
		q = nil
		for i, node := range temp {
			if i+1 < len(temp) {
				node.Next = temp[i+1]
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return root
}
