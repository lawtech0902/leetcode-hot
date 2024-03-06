/*
 * Author: robin-luo
 * Created time: 2024-03-05 14:56:07
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-05 15:08:51
 */

package solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	this.res = strings.Split(data, ",")
	return this.dfs()
}

func (this *Codec) dfs() *TreeNode {
	node := this.res[0]
	this.res = this.res[1:]
	if node == "#" {
		return nil
	}

	val, _ := strconv.Atoi(node)
	root := &TreeNode{
		Val:   val,
		Left:  this.dfs(),
		Right: this.dfs(),
	}

	return root
}
