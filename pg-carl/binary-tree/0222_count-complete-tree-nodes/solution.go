/*
 * Author: robin-luo
 * Created time: 2024-02-22 11:02:09
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 11:09:54
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := 0, 0
	l, r := root.Left, root.Right
	for l != nil {
		l = l.Left
		lh++
	}

	for r != nil {
		r = r.Right
		rh++
	}

	if lh == rh {
		return (2 << lh) - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
