/*
 * Author: robin-luo
 * Created time: 2024-02-22 14:27:47
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-22 14:39:26
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	if len(inorder) == 1 {
		return root
	}

	index := 0
	for i, v := range inorder {
		if v == root.Val {
			index = i
			break
		}
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
