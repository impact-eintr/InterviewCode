package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := []int{}
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		queue = append(queue, root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left.Val)
		}
		if root.Right != nil {
			queue = append(queue, root.Right.Val)
		}
	}

}
