package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	node5 := &ListNode{
		Val:  1,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  2,
		Next: node5,
	}
	node3 := &ListNode{
		Val:  3,
		Next: node4,
	}
	node2 := &ListNode{
		Val:  2,
		Next: node3,
	}
	node1 := &ListNode{
		Val:  1,
		Next: node2,
	}
	cur := node1
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Println()
	res := isPalindrome(node1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func helper(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	fmt.Println(slow.Val)
	return slow

}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	left := head
	right := reverse(helper(left).Next) //后半部分的头

	for right.Next != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	pre := &ListNode{Next: nil}
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
