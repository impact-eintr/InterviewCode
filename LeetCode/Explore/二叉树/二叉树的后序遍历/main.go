package main

import "fmt"

func main() {
	fmt.Println("")

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root == nil {
			return

		}

		traversal(root.Left)
		traversal(root.Right)
		res = append(res, root.Val)
	}
	traversal(root)
	return
}
