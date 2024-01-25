/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 10:10'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	preorder = preorder[1:]
	root := &TreeNode{Val: rootVal}
	index := 0
	for i, val := range inorder {
		if rootVal == val {
			index = i
			break
		}
	}

	root.Left = buildTree(preorder[:index], inorder[:index])
	root.Right = buildTree(preorder[index:], inorder[index+1:])
	return root
}
