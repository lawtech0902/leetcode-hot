/*
__author__ = 'robin-luo'
__date__ = '2023/11/02 09:45'
*/

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	res := false
	rc1 := getGenerator(root1)
	rc2 := getGenerator(root2)
	for {
		rv1, rok1 := <-rc1
		rv2, rok2 := <-rc2
		if !rok1 || !rok2 {
			if rok1 == rok2 {
				res = true
			}

			break
		} else if rv1 != rv2 {
			break
		}
	}

	return res
}

func getGenerator(root *TreeNode) chan int {
	c := make(chan int)
	go func() {
		rangeTree(root, c)
		close(c)
	}()

	return c
}

func rangeTree(root *TreeNode, c chan int) {
	if root.Left == nil && root.Right == nil {
		c <- root.Val
		return
	}

	if root.Left != nil {
		rangeTree(root.Left, c)
	}

	if root.Right != nil {
		rangeTree(root.Right, c)
	}
}
