/*
 * Author: robin-luo
 * Created time: 2024-02-08 15:54:18
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-08 16:16:06
 */

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo = make(map[*TreeNode]int)

// memo + iteration
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	if v, ok := memo[root]; ok {
		return v
	}

	// 偷父节点
	v1 := root.Val
	if root.Left != nil {
		v1 += rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		v1 += rob(root.Right.Left) + rob(root.Right.Right)
	}

	// 不偷父节点
	v2 := rob(root.Left) + rob(root.Right)
	memo[root] = max(v1, v2)
	return memo[root]
}

// dp
func rob1(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

func robTree(cur *TreeNode) [2]int {
	if cur == nil {
		return [2]int{0, 0}
	}

	left := robTree(cur.Left)
	right := robTree(cur.Right)
	// 偷cur，那么就不能偷左右节点
	v1 := cur.Val + left[0] + right[0]
	// 不偷cur，那么可以偷也可以不偷左右节点，则取较大的情况
	v2 := max(left[0], left[1]) + max(right[0], right[1])
	return [2]int{v2, v1}
}
