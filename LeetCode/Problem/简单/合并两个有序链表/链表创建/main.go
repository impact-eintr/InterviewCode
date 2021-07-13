package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	numlist1 := []int{1, 2, 4}
	l1 := createLinkList(numlist1)
	printLinkList(l1) //遍历结果

	numlist2 := []int{1, 1, 3, 4}
	l2 := createLinkList(numlist2)
	printLinkList(l2) //遍历结果

	l3 := mergeTwoLists2(l1, l2)
	printLinkList(l3)

	//	l1 := createLinkList()
}

//归并排序解法
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	head := guard
	for {
		if l1 == nil {
			head.Next = l2
			break
		}
		if l2 == nil {
			head.Next = l1
			break
		}
		if l1.Val < l2.Val {
			v := l1.Val
			for l1 != nil && v == l1.Val {
				node := l1
				l1 = l1.Next
				head.Next = node
				head = head.Next
			}
		} else {
			v := l2.Val
			for l2 != nil && v == l2.Val {
				node := l2
				l2 = l2.Next
				head.Next = node
				head = head.Next
			}
		}
	}
	return guard.Next
}
func createLinkList(numlist []int) *ListNode {
	var head = new(ListNode)
	var tail *ListNode
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	for _, i := range numlist {
		var node = ListNode{Val: i}
		(*tail).Next = &node
		tail = &node
	}
	return head
}

//func insertLinkList(srcList *ListNode, num int) *ListNode {
//	var newNode = ListNode{
//		Val:  num,
//		Next: nil,
//	}
//	temp := srcList
//	for ; temp != nil; temp = temp.Next {
//	}
//	temp.Next = &newNode
//	fmt.Println("in insertfunc ", srcList.Val)
//	return srcList
//}

func printLinkList(srcList *ListNode) {
	for cur := srcList.Next; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Val)
	}
	fmt.Println()
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {
	guard := &ListNode{}
	cur := guard
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			sameVal := l1.Val
			for l1 != nil && sameVal == l1.Val {
				node := l1
				l1 = l1.Next
				cur.Next = node
				cur = cur.Next

			}

		} else {
			sameVal := l2.Val
			for l2 != nil && sameVal == l2.Val {
				//for sameVal == l2.Val && l2 != nil {//首先要有值 然后才能判断是否相等
				node := l2
				l2 = l2.Next
				cur.Next = node
				cur = cur.Next

			}

		}

	}
	if l1 == nil {
		cur.Next = l2

	}
	if l2 == nil {
		cur.Next = l1

	}
	return guard.Next
}
