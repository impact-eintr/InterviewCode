package main

import "fmt"
import "time"

func main() {

	node5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	node4 := &ListNode{
		Val:  4,
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
	test := oddEvenList(node1)
	for test != nil {
		fmt.Print(" ", test.Val)
		test = test.Next
		time.Sleep(time.Second)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	h1 := head
	h2 := head.Next
	H2 := h2
	for h1.Next != nil && h1.Next.Next != nil {
		h1.Next = h1.Next.Next
		h1 = h1.Next
		h2.Next = h2.Next.Next
		h2 = h2.Next
		time.Sleep(time.Second)
	}
	h1.Next = H2
	return head
}
func oddEvenList2(head *ListNode) *ListNode {
	if head == nil {
		return head

	}
	evenHead := head.Next
	odd := head
	even := evenHead
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next

	}
	odd.Next = evenHead
	return head

}
