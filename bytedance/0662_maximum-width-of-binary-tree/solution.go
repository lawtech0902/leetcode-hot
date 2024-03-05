/*
 * Author: robin-luo
 * Created time: 2024-03-04 11:28:11
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 11:37:25
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Pair struct {
	Node  *TreeNode
	Index int
}

func widthOfBinaryTree(root *TreeNode) int {
	res := 1
	q := []Pair{{root, 1}}
	for len(q) != 0 {
		res = max(res, q[len(q)-1].Index-q[0].Index+1)
		temp := q
		q = nil
		for _, pair := range temp {
			if pair.Node.Left != nil {
				q = append(q, Pair{pair.Node.Left, 2 * pair.Index})
			}

			if pair.Node.Right != nil {
				q = append(q, Pair{pair.Node.Right, 2*pair.Index + 1})
			}
		}
	}

	return res
}
