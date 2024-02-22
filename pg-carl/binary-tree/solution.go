/*
 * Author: robin-luo
 * Created time: 2024-02-21 20:17:00
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-21 20:55:10
 */

package solution

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Recursion
func preorderTraversal1(root *TreeNode) []int {
	var (
		res       []int
		traversal func(node *TreeNode)
	)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	var (
		res       []int
		traversal func(root *TreeNode)
	)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}

	traversal(root)
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var (
		res       []int
		traversal func(root *TreeNode)
	)

	traversal = func(root *TreeNode) {
		if root == nil {
			return
		}

		traversal(root.Left)
		res = append(res, root.Val)
		traversal(root.Right)
	}

	traversal(root)
	return res
}

func levelOrderTraversal1(root *TreeNode) [][]int {
	var res [][]int
	var traversal func(root *TreeNode, level int)
	traversal = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(res) == level {
			res = append(res, []int{})
		}

		res[level] = append(res[level], root.Val)
		traversal(root.Left, level+1)
		traversal(root.Right, level+1)
	}

	traversal(root, 0)
	return res
}

// Iteration
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return res
}

func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	reverse(res)
	return res
}

func reverse(a []int) {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var (
		res   []int
		stack []*TreeNode
	)

	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}

	return res
}

func levelOrderTraversal2(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		tempArr := make([]int, len(q))
		for i := range tempArr {
			node := q[0]
			q = q[1:]
			tempArr[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, tempArr)
	}

	return res
}
