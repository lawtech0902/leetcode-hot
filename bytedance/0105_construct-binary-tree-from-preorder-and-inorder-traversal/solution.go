/*
 * Author: robin-luo
 * Created time: 2024-02-29 17:39:36
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-29 17:41:42
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(preorder, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	index := 0
	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(preorder[:index], inorder[:index])
	root.Right = buildTree(preorder[index:], inorder[index+1:])
	return root
}
