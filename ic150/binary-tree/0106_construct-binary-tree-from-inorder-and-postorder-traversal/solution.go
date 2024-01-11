/*
__author__ = 'robin-luo'
__date__ = '2024/01/09 10:05'
*/

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(inorder, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return root
	}

	index := getIndex(root.Val, inorder)
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func getIndex(rootVal int, inorder []int) int {
	for i, num := range inorder {
		if num == rootVal {
			return i
		}
	}

	return 0
}
