package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

var target *ListNode = &ListNode{
	Next: nil,
}

func main() {
	head := ListNode{Next: nil}
	cur := &head
	for i := 22; i < 50; i++ {
		//if cur.Next == nil {
		//	cur.Val = i
		//}
		temp := ListNode{
			Val:  i,
			Next: nil,
		}
		temp.Next = cur.Next
		cur.Next = &temp
	}
	for temp := head.Next; temp.Next != nil; temp = temp.Next {
		fmt.Printf("%d ", temp.Val)
	}
	fmt.Println()
	newhead := reverseBetween(head.Next, 12, 22)
	for temp := newhead.Next; temp != nil; temp = temp.Next {
		fmt.Printf("%d ", temp.Val)
	}
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last

}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		target = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = target
	return last

}
