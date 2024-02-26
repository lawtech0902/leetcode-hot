/*
 * Author: robin-luo
 * Created time: 2024-02-22 09:56:22
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 10:01:51
 */

package solution

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*Node{root}
	for len(q) != 0 {
		curLen := len(q)
		tempArr := make([]int, curLen)
		for i := 0; i < curLen; i++ {
			node := q[0]
			q = q[1:]
			tempArr[i] = node.Val
			q = append(q, node.Children...)
		}

		res = append(res, tempArr)
	}

	return res
}
