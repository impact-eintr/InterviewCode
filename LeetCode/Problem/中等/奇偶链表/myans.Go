package 奇偶链表

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	slow, fast := head.Next, head.Next
	tmp := &ListNode{}
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		tmp.Val = slow.Val
		slow.Val = fast.Val
		fast.Val = tmp.Val
	}
	slower := slow.Next
	for slow.Next != nil && slow.Next.Next != nil {
		slower = slower.Next
		slow = slow.Next.Next

		tmp.Val = slower.Val
		slower.Val = slow.Val
		slow.Val = tmp.Val
	}
	return head

}
