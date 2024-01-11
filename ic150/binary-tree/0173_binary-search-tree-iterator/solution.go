/*
__author__ = 'robin-luo'
__date__ = '2024/01/11 9:48'
*/

package solution

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) BSTIterator {
	var it BSTIterator
	it.inorder(root)
	it.arr = append(it.arr, math.MinInt32)
	return it
}

func (this *BSTIterator) inorder(root *TreeNode) {
	if root == nil {
		return
	}

	this.inorder(root.Left)
	this.arr = append(this.arr, root.Val)
	this.inorder(root.Right)
}

func (this *BSTIterator) Next() int {
	val := this.arr[0]
	this.arr = this.arr[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.arr[0] != math.MinInt32
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
