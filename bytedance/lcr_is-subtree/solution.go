/*
 * Author: robin-luo
 * Created time: 2024-03-04 20:14:50
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-04 20:22:29
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 子结构和子树leetcode 0572不一样，只需要子树的一部分
func isSubStructure(A, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func recur(A, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
