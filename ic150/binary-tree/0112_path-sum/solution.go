/*
__author__ = 'robin-luo'
__date__ = '2024/01/09 19:51'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

type Pair struct {
	Node    *TreeNode
	PathSum int
}

func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []Pair{{root, root.Val}}
	for len(stack) != 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pair.Node.Left == nil && pair.Node.Right == nil {
			return true
		}

		if pair.Node.Right != nil {
			stack = append(stack, Pair{pair.Node.Right, pair.PathSum + pair.Node.Right.Val})
		}

		if pair.Node.Left != nil {
			stack = append(stack, Pair{pair.Node.Left, pair.PathSum + pair.Node.Left.Val})
		}
	}

	return false
}
